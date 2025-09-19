import { FrameworkElement } from 'go-web-framework/framework-element.js';
import { Router } from 'go-web-framework/router.js';

class AppRootAdmin extends FrameworkElement {
  connectedCallback() {
    const router = new Router(this);
    router.baseUrl = '/admin';
    router.addRoute({
      path: '/',
      component: 'manage-users-page',
      importer: () => import('./manage-users-page.js'),
      title: 'Admin Dashboard'
    });
    router.navigate(window.location.pathname);
  }

  disconnectedCallback() {
    console.log('Disconnected ');
  }
}

customElements.define('app-root-admin', AppRootAdmin);
