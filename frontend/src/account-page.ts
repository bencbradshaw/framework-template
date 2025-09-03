import { consume } from 'go-web-framework/context.js';
import { FrameworkElement } from 'go-web-framework/framework-element.js';
import { html } from 'go-web-framework/html.js';
import { reactive } from 'go-web-framework/reactive.js';
import sse from 'go-web-framework/sse.js';
class AccountPage extends FrameworkElement {
  @consume('service') service: any;
  @reactive() message1: string = 'account';
  @reactive() userData: any = null;
  @reactive() apiResponse: string = '';

  connectedCallback() {
    this.updateComplete.then(() => {
      console.log('AccountPage component connected and updated');
      this.getUserData();
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
  }

  getUserData() {
    fetch('/api/user')
      .then((response) => {
        if (response.ok) {
          return response.json();
        }
        return response
          .json()
          .then((err) => {
            throw new Error(JSON.stringify(err));
          })
          .catch(() => {
            throw new Error('Failed to fetch user data');
          });
      })
      .then((userData) => {
        console.log('GET /api/me response:', userData);
        this.userData = userData;
        this.apiResponse = 'User data fetched successfully.';
      })
      .catch((error) => {
        console.error('Error fetching user data (GET):', error);
        if (error instanceof Error) {
          this.apiResponse = error.message;
        } else {
          this.apiResponse = String(error);
        }
      });
  }

  updateUserData() {
    if (!this.userData) {
      this.apiResponse = 'No user data to update.';
      console.error(this.apiResponse);
      return;
    }
    const updatedData = {
      displayName: 'Updated Name ' + new Date().toLocaleTimeString(),
      plan: this.userData.plan === 'premium' ? 'standard' : 'premium'
    };

    fetch('/api/me', {
      method: 'PATCH',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(updatedData)
    })
      .then((response) => {
        if (!response.ok) {
          return response.json().then((err) => {
            throw new Error(JSON.stringify(err));
          });
        }
        return response.json();
      })
      .then((data) => {
        console.log('PATCH /api/me response:', data);
        this.userData = data;
        this.apiResponse = 'User data partially updated successfully.';
      })
      .catch((error) => {
        console.error('Error updating user data (PATCH):', error);
        if (error instanceof Error) {
          this.apiResponse = error.message;
        } else {
          this.apiResponse = String(error);
        }
      });
  }

  replaceUserData() {
    if (!this.userData) {
      this.apiResponse = 'No user data to replace.';
      console.error(this.apiResponse);
      return;
    }

    const newData = {
      ...this.userData,
      displayName: 'Replaced User ' + new Date().toLocaleTimeString()
    };

    fetch('/api/me', {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(newData)
    })
      .then((response) => {
        if (!response.ok) {
          return response.json().then((err) => {
            throw new Error(JSON.stringify(err));
          });
        }
        return response.json();
      })
      .then((data) => {
        console.log('PUT /api/me response:', data);
        this.userData = data;
        this.apiResponse = 'User data fully updated successfully.';
      })
      .catch((error) => {
        console.error('Error replacing user data (PUT):', error);
        if (error instanceof Error) {
          this.apiResponse = error.message;
        } else {
          this.apiResponse = String(error);
        }
      });
  }

  get styles() {
    return html`
      <style>
        /* --- Main Content --- */
        .container {
          max-width: 1200px;
          margin: 0 auto;
          padding: 3rem 2rem;
        }

        .page-header h1 {
          font-size: clamp(2rem, 6vw, 2.5rem);
          font-weight: 700;
          color: #fff;
          margin-bottom: 2rem;
        }

        /* --- Account Layout --- */
        .account-layout {
          display: grid;
          grid-template-columns: 240px 1fr;
          gap: 3rem;
        }

        /* --- Account Navigation Sidebar --- */
        .account-nav {
          display: flex;
          flex-direction: column;
        }

        .account-nav a {
          color: #94a3b8;
          text-decoration: none;
          padding: 0.75rem 1rem;
          border-radius: 8px;
          margin-bottom: 0.5rem;
          transition: background-color 0.2s ease, color 0.2s ease;
        }

        .account-nav a:hover {
          background-color: #1f2937;
          color: #fff;
        }

        .account-nav a.active {
          background-color: #4f46e5;
          color: #fff;
          font-weight: bold;
        }

        /* --- Account Content Area --- */
        .account-content .card {
          background: #111827;
          border: 1px solid #1f2937;
          border-radius: 12px;
          padding: 2rem;
          margin-bottom: 2rem;
        }

        .card-header {
          display: flex;
          justify-content: space-between;
          align-items: center;
          margin-bottom: 1.5rem;
          padding-bottom: 1.5rem;
          border-bottom: 1px solid #1f2937;
        }

        .card-header h2 {
          font-size: 1.5rem;
          color: #fff;
        }

        .edit-link {
          font-size: 0.9rem;
          color: #4f46e5;
          text-decoration: none;
          font-weight: bold;
        }

        .info-row {
          display: flex;
          justify-content: space-between;
          align-items: center;
          margin-bottom: 1rem;
        }

        .info-row .label {
          color: #94a3b8;
        }

        .info-row .value {
          color: #e2e8f0;
          font-family: 'Roboto Mono', monospace;
        }

        .subscription-plan {
          background-color: #4f46e5;
          color: #fff;
          padding: 0.25rem 0.75rem;
          border-radius: 20px;
          font-size: 0.9rem;
          font-weight: bold;
        }

        .cta-button {
          display: inline-block;
          padding: 0.75rem 1.5rem;
          font-size: 0.9rem;
          font-weight: bold;
          color: #fff;
          background-color: #374151;
          border: none;
          border-radius: 8px;
          cursor: pointer;
          text-decoration: none;
          transition: background-color 0.3s;
        }

        .cta-button:hover {
          background-color: #4b5563;
        }

        /* --- Responsive --- */
        @media (max-width: 768px) {
          .account-layout {
            grid-template-columns: 1fr;
          }
        }

        /* --- Footer --- */
        .footer {
          text-align: center;
          padding: 2rem;
          margin-top: 4rem;
          border-top: 1px solid #1f2937;
          color: #64748b;
          font-size: 0.9rem;
        }
      </style>
    `;
  }

  render() {
    return html`
      ${this.styles}
      <div class="container">
        <div class="account-layout">
          <aside class="account-nav">
            <a href="#" class="active">Account Overview</a>
            <a href="#">Security & Password</a>
            <a href="/account/manage-plan">Manage Plan</a>
            <a href="/logout" router-ignore id="sign-out" style="margin-top: 2rem; color: #ef4444; padding: 1rem;"
              >Sign Out</a
            ></a>
          </aside>

          <main class="account-content">
            <header class="page-header">
              <h1>Account Overview</h1>
            </header>


            ${
              this.userData
                ? html`
                    <div class="card">
                      <div class="card-header">
                        <h2>User Profile</h2>
                        <a href="#" class="edit-link">Edit</a>
                      </div>
                      <div class="info-row">
                        <span class="label">Display Name</span>
                        <span class="value">${this.userData.displayName}</span>
                      </div>
                      <div class="info-row">
                        <span class="label">Email</span>
                        <span class="value">${this.userData.email}</span>
                      </div>
                    </div>

                    <div class="card">
                      <div class="card-header">
                        <h2>Subscription Details</h2>
                      </div>
                      <div class="info-row">
                        <span class="label">Current Plan</span>
                        <span class="value">${this.userData.plan ? this.userData.plan : 'Free'}</span>
                      </div>
                      <div class="info-row">
                        <span class="label">Status</span>
                        <span class="value subscription-plan">Active</span>
                      </div>
                    </div>
                  `
                : html`<p>No user data found. It might have been deleted or is still loading.</p>`
            }
          </main>
        </div>
      </div>
    `;
  }
}

customElements.define('account-page', AccountPage);
