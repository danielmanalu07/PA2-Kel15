// import 'package:flutter/material.dart';
// import 'package:get/get.dart';
// import 'package:the_deck/Models/Product.dart';

// class ProductListController extends GetxController {
//   late List<Products> _products = [];

//   List<Products> get products => _products;

//   void fetchProducts(String categoryId) {
//     // Fetch products based on the category ID
//     _products = getProductsByCategory(categoryId);
//     update(); // Perbarui widget ketika daftar produk berubah
//   }
// }

// class ProductListScreen extends StatelessWidget {
//   final String categoryId;
//   final ProductListController productListController =
//       Get.put(ProductListController());

//   const ProductListScreen({required this.categoryId});

//   @override
//   Widget build(BuildContext context) {
//     productListController.fetchProducts(categoryId);

//     return Scaffold(
//       appBar: AppBar(
//         title: Text('Products'),
//       ),
//       body: Obx(() {
//         // Gunakan Obx di sini untuk memperbarui widget saat daftar produk berubah
//         return ListView.builder(
//           itemCount: productListController.products.length,
//           itemBuilder: (context, index) {
//             final product = productListController.products[index];
//             return ListTile(
//               title: Text(product.name),
//             );
//           },
//         );
//       }),
//     );
//   }
// }

// class Products {
//   final String id;
//   final String name;
//   final String image;
//   final String description;
//   final double price;
//   final String categoryID;

//   Products({
//     required this.id,
//     required this.name,
//     required this.image,
//     required this.price,
//     required this.description,
//     required this.categoryID,
//   });

//   factory Products.fromJson(Map<String, dynamic> json) {
//     return Products(
//       id: json['id'],
//       name: json['name'],
//       image: json['image'],
//       price: double.parse(json['price']),
//       description: json['description'],
//       categoryID: json['category_id'],
//     );
//   }
// }

// List<Products> getProductsByCategory(String categoryId) {
//   // Mock data
//   return [
//     Products(id: '1', name: 'Product 1', image: 'image1.jpg', price: 10.0, description: 'Description 1', categoryID: 'category1'),
//     Products(id: '2', name: 'Product 2', image: 'image2.jpg', price: 20.0, description: 'Description 2', categoryID: 'category1'),
//     // Add more products here
//   ];
// }
