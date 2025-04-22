import { provide } from 'go-web-framework/context.js';
import { FrameworkElement } from 'go-web-framework/framework-element.js';
import { Router } from 'go-web-framework/router.js';
import sse from 'go-web-framework/sse.js';
import { Store } from 'src/store.js';

class AppRoot extends FrameworkElement {
  @provide('service') service = [1, 2, 3, 4];
  @provide('store') store = new Store();
  unsubscribe: () => void;

  constructor() {
    super();
    this.unsubscribe = sse('/events', (event: string, data: any) => {
      console.log('SSE event received:', event, data);
      if (event === 'entity' && 'user' in data) {
        this.store.users = [...this.store.users, data.user];
      }
      if (event === 'esbuild') {
        console.log('esbuild event received');
        window.location.reload();
      }
    });
  }

  connectedCallback() {
    const router = new Router(this);
    router.baseUrl = '/account';
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

  disconnectedCallback() {
    console.log('Disconnected');
    this.unsubscribe();
  }
}

customElements.define('app-root', AppRoot);
