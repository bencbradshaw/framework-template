import { consume } from 'go-web-framework/context.js';
import { FrameworkElement } from 'go-web-framework/framework-element.js';
import { html } from 'go-web-framework/html.js';
import { reactive } from 'go-web-framework/reactive.js';
import { Store } from 'src/store.js';
import './app-nav.js';

class User extends FrameworkElement {
  @consume('store') store: Store;

  @reactive() instanceId: string = Math.random().toString().slice(2, 10);
  @reactive() users: { ID: number; Email: string }[] = [];

  async connectedCallback() {
    console.log('store', this.store);
    this.store.subscribe('users', (users) => {
      console.log('users SUB', users);
      this.users = users;
    });

    this.updateComplete.then(() => this.attachButtonListener());
  }

  update() {
    super.update();
    this.updateComplete.then(() => this.attachButtonListener());
  }

  attachButtonListener() {
    const button = this.shadowRoot?.querySelector('button');
    if (button) {
      button.addEventListener('click', async () => {
        console.log('click');
        const resp = await fetch('/api/v1/user', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({
            first_name: this.shadowRoot?.querySelector<HTMLInputElement>('#firstname')?.value,
            last_name: this.shadowRoot?.querySelector<HTMLInputElement>('#lastname')?.value,
            email: this.shadowRoot?.querySelector<HTMLInputElement>('#email')?.value
          })
        });
        console.log('resp', await resp.json());
      });
    }
  }

  render() {
    return html`
      <style>
        :host {
          display: flex;
          flex-direction: column;
          align-items: center;
          justify-content: center;
        }
        section {
          display: flex;
          flex-direction: column;
          padding: 1rem;
        }
      </style>
      <app-nav></app-nav>
      <section>
        <label>First Name</label>
        <input type="text" id="firstname" />
        <label>Last Name</label>
        <input type="text" id="lastname" />
        <label> email </label>
        <input type="email" id="email" />
        <button>Submit</button>
      </section>
      <section>
        <h2>Users</h2>
        <ul>
          ${this.users
            ?.toReversed()
            .map((user) => html` <li>${user.Email}</li> `)
            .join('')}
        </ul>
      </section>
    `;
  }
}

customElements.define('app-user', User);
