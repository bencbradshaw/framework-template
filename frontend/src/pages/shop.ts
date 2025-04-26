import { FrameworkElement, html, reactive } from 'go-web-framework';

class Shop extends FrameworkElement {
  @reactive() items: { id: number; name: string }[] = [];

  connectedCallback() {
    this.updateComplete.then(() => {
      this.attachButtonListener();
    });
  }

  attachButtonListener() {
    const button = this.shadowRoot?.querySelector('button');
    if (button) {
      button.addEventListener('click', async () => {
        const resp = await fetch('/api/shop');
        this.items = [...this.items, ...(await resp.json())];
      });
    }
  }

  render() {
    return html`
      <div>
        <h1>Shop</h1>
        <ul>
          ${this.items.map((item) => `<li>${item.name}</li>`).join('')}
        </ul>
        <button>Show More</button>
      </div>
    `;
  }
}
