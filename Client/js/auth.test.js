const { JSDOM } = require('jsdom');
const { showUserMenu } = require('./auth.js');

// Mock the fetch function
global.fetch = jest.fn(() =>
  Promise.resolve({
    json: () => Promise.resolve({}),
  })
);

describe('showUserMenu', () => {
  beforeEach(() => {
    const dom = new JSDOM(`<!DOCTYPE html><html><body><header><button class="auth-btn">Login</button></header></body></html>`);
    global.document = dom.window.document;
    global.window = dom.window;
    // Ensure the header is directly in the document body for testing purposes
    global.document.body.innerHTML = `<header><button class="auth-btn">Login</button></header>`;
  });

  test('should remove the auth button and add the profile menu', () => {
    const user = { email: 'test@example.com', avatar: 'avatar.png' };
    showUserMenu(user);

    const authBtn = document.querySelector('.auth-btn');
    expect(authBtn).toBeNull();

    const profileMenu = document.querySelector('.profile-menu');
    expect(profileMenu).not.toBeNull();

    const avatar = document.querySelector('.avatar-icon');
    expect(avatar.src).toContain('avatar.png');
  });
});
