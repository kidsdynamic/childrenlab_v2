import { SitePage } from './app.po';

describe('site App', () => {
  let page: SitePage;

  beforeEach(() => {
    page = new SitePage();
  });

  it('should display message saying app works', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('app works!');
  });
});
