import { provide } from 'go-web-framework/context.js';
import { FrameworkElement } from 'go-web-framework/framework-element.js';
import { Router } from 'go-web-framework/router.js';
import { Store } from 'src/store.js';

class AppRoot extends FrameworkElement {
  @provide('service') service = [1, 2, 3, 4];
  @provide('store') store = new Store();

  constructor() {
    super();
  }

  connectedCallback() {
    const router = new Router(this);
    router.baseUrl = '/app';
    router.addRoute({
      path: '/',
      component: 'app-landing',
      importer: () => import('../spa/app-landing.js'),
      title: 'App Landing 1'
    });
    router.addRoute({
      path: '/user',
      component: 'app-user',
      importer: () => import('../spa/app-user.js'),
      title: 'App User'
    });
    // Navigate to the initial route
    router.navigate(window.location.pathname);
  }
}

customElements.define('app-root', AppRoot);
