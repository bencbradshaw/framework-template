import { FrameworkElement } from 'go-web-framework/framework-element.js';
import { html } from 'go-web-framework/html.js';
import { reactive } from 'go-web-framework/reactive.js';
import sse from 'go-web-framework/sse.js';
import styles from './manage-users-page.css.js';

interface User {
  id: string;
  email: string;
  name: string;
  role: string;
  created_at: string;
}

class ManageUsersPage extends FrameworkElement {
  @reactive() currentUser: User | null = null;
  @reactive() allUsers: User[] = [];
  connectedCallback() {
    this.updateComplete.then(() => {
      console.log('ManageUsersPage component connected and updated');
      this.getCurrentUser();
      this.getAllUsers();
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
    // Use event delegation for dynamic content
    const container = this.shadowRoot.querySelector('.users-table');
    if (container) {
      container.addEventListener('click', this.#handleTableClick);
    }
  }

  disconnectedCallback() {
    // Remove event listeners
    const container = this.shadowRoot.querySelector('.users-table');
    if (container) {
      container.removeEventListener('click', this.#handleTableClick);
    }
  }

  #handleTableClick = (event: Event) => {
    const target = event.target as HTMLElement;
    const row = target.closest('.table-row') as HTMLElement;

    if (!row) return;

    const userId = row.getAttribute('data-user-id');
    if (!userId) return;

    if (target.classList.contains('btn-view')) {
      this.viewUser(userId);
    } else if (target.classList.contains('btn-delete')) {
      this.deleteUser(userId);
    }
  };

  async getCurrentUser() {
    const { data } = (await fetch('/api/user').then((r) => r.json())) as { success: boolean; data: User };
    this.currentUser = data;
  }

  async getAllUsers() {
    try {
      const { data } = (await fetch('/api/users').then((r) => r.json())) as { success: boolean; data: User[] };
      this.allUsers = data;
      console.log('Fetched users:', this.allUsers);
    } catch (error) {
      console.error('Failed to fetch users:', error);
      this.allUsers = [];
    }
  }

  render() {
    return html`
      ${styles}
      <div class="container">
        <div class="account-layout">
          <aside class="account-nav">
            <a href="#" class="active">Admin Dashboard</a>
            <a href="/logout" router-ignore id="sign-out" style="margin-top: 2rem; color: #ef4444; padding: 1rem;">
              Sign out
            </a>
          </aside>
          <main class="account-content">
            <header class="page-header">
              <h1>User Management</h1>
              <p>Manage all users in the system</p>
            </header>

            ${this.currentUser
              ? html`
                  <div class="card">
                    <div class="card-header">
                      <h2>Current Admin</h2>
                    </div>
                    <div class="info-row">
                      <span class="label">Email</span>
                      <span class="value">${this.currentUser.email}</span>
                    </div>
                    <div class="info-row">
                      <span class="label">Role</span>
                      <span class="value">${this.currentUser.role}</span>
                    </div>
                  </div>
                `
              : html`<p>Loading admin profile...</p>`}

            <div class="card">
              <div class="card-header">
                <h2>All Users (${this.allUsers?.length})</h2>
              </div>

              ${this.allUsers?.length > 0
                ? html`
                    <div class="users-table">
                      <div class="table-header">
                        <div class="table-cell">Email</div>
                        <div class="table-cell">Role</div>
                        <div class="table-cell">Created</div>
                        <div class="table-cell">Actions</div>
                      </div>
                      ${this.allUsers
                        ?.map(
                          (user) => html`
                            <div class="table-row" data-user-id="${user.id}">
                              <div class="table-cell">
                                <strong>${user.email}</strong>
                              </div>
                              <div class="table-cell">
                                <span class="role-badge role-${user.role}">${user.role}</span>
                              </div>
                              <div class="table-cell">${new Date(user.created_at).toLocaleDateString()}</div>
                              <div class="table-cell">
                                <button class="btn-small btn-view">View</button>
                                ${user.role !== 'admin'
                                  ? html`<button class="btn-small btn-danger btn-delete">Delete</button>`
                                  : html`<span class="disabled">Protected</span>`}
                              </div>
                            </div>
                          `
                        )
                        .join('')}
                    </div>
                  `
                : html`<p>No users found or still loading...</p>`}
            </div>
          </main>
        </div>
      </div>
    `;
  }

  viewUser(userId: string) {
    console.log('View user:', userId);
    // Implement user details view
  }

  deleteUser(userId: string) {
    if (confirm('Are you sure you want to delete this user?')) {
      console.log('Delete user:', userId);
      // Implement user deletion
    }
  }
}

customElements.define('manage-users-page', ManageUsersPage);
