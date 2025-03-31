import { consume } from 'go-web-framework/context.js';
import { FrameworkElement } from 'go-web-framework/framework-element.js';
import { html } from 'go-web-framework/html.js';
import { reactive } from 'go-web-framework/reactive.js';
import './app-nav.js';
class Landing extends FrameworkElement {
  @consume('service') service: any;
  @reactive() message1: string = 'landing';

  connectedCallback() {
    console.log('service', this.service);
  }

  render() {
    return html`<app-nav></app-nav>
      <p>${this.message1}</p>`;
  }
}

customElements.define('app-landing', Landing);
