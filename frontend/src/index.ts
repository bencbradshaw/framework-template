// import the index.css so esbuild will output it with the bundle
import './index.css';

// split apart some js files individually for specific routes
window.addEventListener('DOMContentLoaded', async () => {
  const path = '/' + (window.location.pathname.split('/')[1] ?? '');
  switch (true) {
    case path === '/app' || /^\/app\/.*/.test(path):
      await import('./app/app-root-user.js');
      break;
    case path === '/admin' || /^\/admin\/.*/.test(path):
      await import('./admin/app-root-admin.js');
      break;
  }
});
