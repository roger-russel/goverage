const puppeteer = require('puppeteer');
const select = require('puppeteer-select');


async function run() {
  const browser = await puppeteer.launch();
  const page = await browser.newPage();

  await page.goto('file://'+process.cwd()+'/tmp/coverage.html');

  const element = await select(page).getElement('.table-list table tbody tr');
  await element.click()

  await page.screenshot({ path: './tmp/coverage.png', fullPage: true });

  browser.close();
}

run();
