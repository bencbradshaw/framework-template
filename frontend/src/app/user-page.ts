import { FrameworkElement } from 'go-web-framework/framework-element.js';
import { html } from 'go-web-framework/html.js';
import { reactive } from 'go-web-framework/reactive.js';
import sse from 'go-web-framework/sse.js';
import styles from './user-page.css.js';

interface User {
  email: string;
  name: string;
  plan?: string;
}

class UserPage extends FrameworkElement {
  @reactive() user: any = null;
  connectedCallback() {
    this.updateComplete.then(() => {
      console.log('UserPage component connected and updated');
      this.getUser();
    });
    sse('/events', (event: string, data: any) => {
      console.log('SSE event received:', event, data);
      if (event === 'esbuild') {
        console.log('esbuild event received');
        window.location.reload();
      }
    });
  }

  update(): void {
    super.update();
    // if needing to add event listeners, do it here, so
    // after each lifecycle update, the listeners are reattached
    this.shadowRoot.querySelector('#status')?.addEventListener('click', this.#handleActiveClick);
  }

  disconnectedCallback() {
    this.shadowRoot.querySelector('#status')?.removeEventListener('click', this.#handleActiveClick);
  }

  #handleActiveClick = () => {
    console.log('Active plan clicked');
  };

  async getUser() {
    const { data } = (await fetch('/api/user').then((r) => r.json())) as { success: boolean; data: User };
    this.user = data;
  }

  render() {
    return html`
      ${styles}
      <div class="container">
        <div class="account-layout">
          <aside class="account-nav">
            <a href="#" class="active">Account Overview</a>
            <a href="/logout" router-ignore id="sign-out" style="margin-top: 2rem; color: #ef4444; padding: 1rem;">
              Sign out
            </a>
          </aside>
          <main class="account-content">
            <header class="page-header">
              <h1>Account Overview</h1>
            </header>
            ${this.user
              ? html`
                  <div class="card">
                    <div class="card-header">
                      <h2>User Profile</h2>
                    </div>
                    <div class="info-row">
                      <span class="label">Display Name</span>
                      <span class="value">${this.user.name}</span>
                    </div>
                    <div class="info-row">
                      <span class="label">Email</span>
                      <span class="value">${this.user.email}</span>
                    </div>
                  </div>

                  <div class="card">
                    <div class="card-header">
                      <h2>Subscription Details</h2>
                    </div>
                    <div class="info-row">
                      <span class="label">Current Plan</span>
                      <span class="value">${this.user.plan ? this.user.plan : 'Free'}</span>
                    </div>
                    <div class="info-row">
                      <span class="label">Status</span>
                      <span id="status" class="value subscription-plan">Active</span>
                    </div>
                  </div>
                `
              : html`<p>No user data found. It might have been deleted or is still loading.</p>`}
          </main>
        </div>
      </div>
    `;
  }
}

customElements.define('user-page', UserPage);
