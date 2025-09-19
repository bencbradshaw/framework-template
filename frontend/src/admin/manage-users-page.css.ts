import { html } from 'go-web-framework/html.js';

export default html`
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
      cursor: pointer;
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
