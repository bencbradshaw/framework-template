import { FrameworkElement } from 'go-web-framework/framework-element.js';
import { Router } from 'go-web-framework/router.js';

class AppRoot extends FrameworkElement {
  connectedCallback() {
    const router = new Router(this);
    router.baseUrl = '/account';
    router.addRoute({
      path: '/',
      component: 'account-page',
      importer: () => import('./account-page.js'),
      title: 'Account Overview'
    });
    router.addRoute({
      path: '/manage-plan',
      component: 'manage-plan-page',
      importer: () => import('./manage-plan-page.js'),
      title: 'Manage Plan'
    });
    router.navigate(window.location.pathname);
  }

  disconnectedCallback() {
    console.log('Disconnected ');
  }
}

customElements.define('app-root', AppRoot);
