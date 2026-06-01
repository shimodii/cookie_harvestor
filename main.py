import asyncio
import json
from pathlib import Path
from playwright.async_api import async_playwright

URL = "https://divar.ir/s/tehran"
AUTH_ENDPOINT = "https://api.divar.ir/v8/auth/open-initiate-page"
TARGET_COOKIES = {"sAccessToken", "sFrontToken"}
OUTPUT_DIR = Path("output")


async def harvest_cookies():
    phone_number = None
    OUTPUT_DIR.mkdir(exist_ok=True)

    async with async_playwright() as p:
        browser = await p.chromium.launch(headless=False)
        context = await browser.new_context()
        page = await context.new_page()

        async def handle_request(request):
            nonlocal phone_number
            if AUTH_ENDPOINT in request.url and request.method == "POST":
                try:
                    body = request.post_data_json
                    phone_number = body["specification"]["data"]["data"]["phone"]["str"]["value"]
                    print(f"Phone number captured: {phone_number}")
                except Exception as e:
                    print(f"Could not extract phone number: {e}")

        page.on("request", handle_request)

        await page.goto(URL)
        print("Browser opened.")
        print("Click the menu -> Login, enter your phone + OTP in the popup.")
        print("Waiting for you to finish logging in...\n")

        for _ in range(90):
            cookies = await context.cookies()
            cookie_names = {c["name"] for c in cookies}

            if TARGET_COOKIES.issubset(cookie_names):
                print("Login detected!")
                break

            await asyncio.sleep(2)
        else:
            print("Timed out. Grabbing whatever cookies exist...")
            cookies = await context.cookies()

        tokens = {c["name"]: c["value"] for c in cookies if c["name"] in TARGET_COOKIES}

        if tokens:
            filename = OUTPUT_DIR / f"{phone_number}.txt" if phone_number else OUTPUT_DIR / "tokens.txt"
            with open(filename, "w") as f:
                json.dump(tokens, f, indent=2)
            print(f"Tokens saved to {filename}")
            for name, value in tokens.items():
                print(f"  {name}: {value}")
        else:
            print("Target cookies not found.")

        await browser.close()
        return tokens


if __name__ == "__main__":
    asyncio.run(harvest_cookies())