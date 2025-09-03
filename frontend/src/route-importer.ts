export async function routeImporter(path: string) {
  console.log('Route importer initialized for path:', path);
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
    case path === '/login':
      await import('./login.js');
      break;
    case path === '/account' || /^\/account\/.*/.test(path):
      await import('./login.js');
      await import('./app-root.js');
      break;
    default:
      await import('./pages/home.js');
  }
}
