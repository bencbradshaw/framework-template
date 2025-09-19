import { FrameworkElement } from 'go-web-framework/framework-element.js';
import { Router } from 'go-web-framework/router.js';

class AppRootUser extends FrameworkElement {
  connectedCallback() {
    const router = new Router(this);
    router.baseUrl = '/app';
    router.addRoute({
      path: '/',
      component: 'user-page',
      importer: () => import('./user-page.js'),
      title: 'User Dashboard'
    });
    router.navigate(window.location.pathname);
  }

  disconnectedCallback() {
    console.log('Disconnected ');
  }
}

customElements.define('app-root-user', AppRootUser);
