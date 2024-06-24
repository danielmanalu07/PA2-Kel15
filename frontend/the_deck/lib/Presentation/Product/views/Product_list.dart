import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:the_deck/Controller/ProductController.dart';
import 'package:the_deck/Models/Product.dart';
import 'package:the_deck/Presentation/Foods/Views/about_menu_view.dart';

class ProductListScreen extends StatelessWidget {
  final int categoryId;
  final ProductController productListController = Get.put(ProductController());

  ProductListScreen({required this.categoryId});

  @override
  Widget build(BuildContext context) {
    productListController.getProductByCategoryId(categoryId);

    return Scaffold(
      appBar: AppBar(
        title: Text('Products'),
      ),
      body: Obx(() {
        if (productListController.productList.isEmpty) {
          return Center(child: Text('No products found.'));
        }
        return ListView.builder(
          itemCount: productListController.productList.length,
          itemBuilder: (context, index) {
            final product = productListController.productList[index];
            return ProductListItem(product: product);
          },
        );
      }),
    );
  }
}

class ProductListItem extends StatelessWidget {
  final Product product;

  const ProductListItem({
    Key? key,
    required this.product,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ListTile(
      title: Text(product.name),
      subtitle: Text('Rp ${product.price.toStringAsFixed(2)}'), // Harga produk
      leading: Container(
        width: 60,
        height: 60,
        decoration: BoxDecoration(
          shape: BoxShape.circle,
          image: DecorationImage(
            image: NetworkImage(
                "http://192.168.188.215:8080/product/image/${product.image}"), // Gambar produk
            fit: BoxFit.cover,
          ),
        ),
      ),
      onTap: () {
        Get.to(() => AboutMenuView(productId: product.id));
      },
    );
  }
}
