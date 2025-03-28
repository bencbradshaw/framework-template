import './index.css';
import { routeImporter } from './route-importer.js';

function main() {
  window.addEventListener('DOMContentLoaded', async () => {
    await routeImporter('/' + (window.location.pathname.split('/')[1] ?? ''));
  });
}

main();
