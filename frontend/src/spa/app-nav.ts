import { FrameworkElement } from 'go-web-framework/framework-element.js';
import { html } from 'go-web-framework/html.js';
class Nav extends FrameworkElement {
  render() {
    return html`<nav>
      <a href="/account/">App</a>
      <a href="/account/user">User</a>
    </nav>`;
  }
}

customElements.define('app-nav', Nav);
