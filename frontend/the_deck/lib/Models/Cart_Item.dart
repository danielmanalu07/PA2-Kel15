class CartItem {
  final int id;
  final int productId;

  CartItem({required this.id, required this.productId});

  factory CartItem.fromJson(Map<String, dynamic> json) {
    return CartItem(
      id: json['id'],
      productId: json['product_id'],
    );
  }
}
