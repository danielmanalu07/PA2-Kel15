import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:the_deck/Controller/ProductController.dart';
import 'package:the_deck/Models/Product.dart';

class ProductListScreen extends StatelessWidget {
  final int categoryId;
  final ProductController productListController = Get.put(ProductController());

  ProductListScreen({required this.categoryId});

  @override
  Widget build(BuildContext context) {
    productListController.GetProductByCategoryId(categoryId);

    return Scaffold(
      appBar: AppBar(
        title: Text('Products'),
      ),
      body: Obx(() {
        if (productListController.products.isEmpty) {
          return Center(child: Text('No products found.'));
        }
        return ListView.builder(
          itemCount: productListController.products.length,
          itemBuilder: (context, index) {
            final product = productListController.products[index];
            return ListTile(
              title: Text(product.name),
            );
          },
        );
      }),
    );
  }
}
