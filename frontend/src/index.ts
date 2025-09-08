// import the index.css so esbuild will output it with the bundle
import './index.css';

// split apart some js files individually for specific routes
// this way, the home page vs the login page, don't load and execute all the javascript
window.addEventListener('DOMContentLoaded', async () => {
  const path = '/' + (window.location.pathname.split('/')[1] ?? '');
  switch (true) {
    case path === '/login':
      await import('./login.js');
      break;
    case path === '/account' || /^\/account\/.*/.test(path):
      await import('./login.js');
      await import('./account/app-root.js');
      break;
  }
});
