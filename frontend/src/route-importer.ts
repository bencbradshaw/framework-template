export async function routeImporter(path: string) {
  switch (true) {
    case path === '/':
      await import('./pages/home.js');
      break;
    case path === '/about':
      await import('./pages/about.js');
      break;
    case path === '/app' || /^\/app\/.*/.test(path):
      await import('./pages/app.js');
      break;
    default:
      await import('./pages/home.js');
  }
}
