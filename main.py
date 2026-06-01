import asyncio
import json
from playwright.async_api import async_playwright

URL = "https://divar.ir/s/tehran"


async def harvest_cookies():
    async with async_playwright() as p:
        browser = await p.chromium.launch(headless=False)
        context = await browser.new_context()
        page = await context.new_page()

        await page.goto(URL)
        print("Browser opened.")
        print("Waiting for you to finish logging in...\n")

        for _ in range(90):
            cookies = await context.cookies()
            cookie_names = [c["name"] for c in cookies]

            if "sAccessToken" in cookie_names:
                print("Login detected!")
                break

            await asyncio.sleep(2)
        else:
            print("Timed out waiting for login. Grabbing whatever cookies exist...")
            cookies = await context.cookies()

        print(f"\nExtracted {len(cookies)} cookies:")
        for c in cookies:
            print(f"  {c['name']}: {c['value']}")

        with open("cookies.json", "w") as f:
            json.dump(cookies, f, indent=2)
        print("\nCookies saved to cookies.json")

        await browser.close()
        return cookies


if __name__ == "__main__":
    asyncio.run(harvest_cookies())