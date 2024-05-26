class CartItem {
  final int id;
  final int productId;
  final int quantity;
  bool isChecked;

  CartItem({
    required this.id,
    required this.productId,
    required this.quantity,
    this.isChecked = false,
  });

  factory CartItem.fromJson(Map<String, dynamic> json) {
    return CartItem(
      id: json['id'],
      productId: json['product_id'],
      quantity: json['quantity'],
    );
  }
}
