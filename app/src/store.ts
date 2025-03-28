import { StateStore, prop } from 'go-web-framework/state-store.js';

export class EntityStore extends StateStore {
  @prop() users: { ID: number; Email: string }[] = [];
}
