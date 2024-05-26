class Product {
  final int id;
  final String name;
  final String image;
  final String description;
  final double price;
  final int CategoryID;

  Product({
    required this.id,
    required this.name,
    required this.image,
    required this.price,
    required this.description,
    required this.CategoryID,
  });

  factory Product.fromJson(Map<String, dynamic> json) {
    return Product(
      id: json['id'],
      name: json['name'],
      image: json['image'],
      price: double.parse(json['price']),
      description: json['description'],
      CategoryID: json['category_id'],
    );
  }
}
