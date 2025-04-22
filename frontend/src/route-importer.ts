export async function routeImporter(path: string) {
  switch (true) {
    case path === '/':
      await import('./pages/home.js');
      break;
    case path === '/shop':
      await import('./pages/shop.js');
      break;
    case path === '/about':
      await import('./pages/about.js');
      break;
    case path === '/account' || /^\/account\/.*/.test(path):
      await import('./pages/account.js');
      break;
    default:
      await import('./pages/home.js');
  }
}
