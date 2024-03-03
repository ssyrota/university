export class Sell {
  /**
   * @param {{
   * id: string,
   * time: string,
   * product_id: string,
   * product_name: string,
   * product_type: string,
   * store_id: string,
   * store_name: string,
   * user_id: string,
   * quantity: number
   * }} data
   */
  constructor(data) {
    this.id = data.id;
    this.time = data.time;
    this.product_id = data.product_id;
    this.product_name = data.product_name;
    this.product_type = data.product_type;
    this.store_id = data.store_id;
    this.store_name = data.store_name;
    this.user_id = data.user_id;
    this.quantity = data.quantity;
  }
}
