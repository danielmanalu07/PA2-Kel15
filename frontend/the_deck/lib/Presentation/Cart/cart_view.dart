import 'package:get/get.dart';
import 'package:the_deck/Controller/CustomerController.dart';
import 'package:the_deck/Controller/ProductController.dart';
import 'package:the_deck/Core/app_colors.dart';
import 'package:the_deck/Core/font_size.dart';
import 'package:the_deck/Core/response_conf.dart';
import 'package:the_deck/Core/text_styles.dart';
import 'package:the_deck/Models/Cart_Item.dart';
import 'package:the_deck/Models/Product.dart';
import 'package:the_deck/Presentation/Auth/screens/default_button.dart';
import 'package:the_deck/Presentation/Base/base.dart';
import 'package:the_deck/Presentation/Base/food_item.dart';
import 'package:the_deck/Presentation/Cart/screens/cart_item_food.dart';
import 'package:flutter/material.dart';
import 'package:gap/gap.dart';
import 'package:the_deck/Presentation/Models/category_model.dart';

class CartView extends StatefulWidget {
  CartView({Key? key}) : super(key: key);

  @override
  State<CartView> createState() => _CartViewState();
}

class _CartViewState extends State<CartView> {
  final RegisterController customterController = Get.put(RegisterController());
  final ProductController _productController = Get.put(ProductController());
  late Future<List<Product>> _productFuture;

  int _totalItems = 0;
  double _totalPrice = 0.0;

  @override
  void initState() {
    super.initState();
    _productFuture = _productController.getProductList();
  }

  void _updatePaymentSummary(List<Product> products) {
    int totalItems = 0;
    double totalPrice = 0.0;
    for (var item in customterController.cartItems) {
      final product =
          products.firstWhere((product) => product.id == item.productId);
      if (item.isChecked) {
        totalItems += item.quantity;
        totalPrice += item.quantity * product.price;
      }
    }
    setState(() {
      _totalItems = totalItems;
      _totalPrice = totalPrice;
    });
  }

  void _updateQuantity(int cartItemId, int newQuantity) async {
    await customterController.updateCartItemQuantity(cartItemId, newQuantity);
    _updatePaymentSummary(await _productFuture);
  }

  @override
  Widget build(BuildContext context) {
    customterController.getMyCart();
    return Scaffold(
      appBar: buildAppBar(
          buildContext: context, screenTitle: "My Cart", isBackup: false),
      body: Padding(
        padding: EdgeInsets.symmetric(horizontal: getWidth(24)),
        child: SingleChildScrollView(
          child: Column(
            children: [
              Obx(() {
                if (customterController.cartItems.isEmpty) {
                  return Text("Cart Item Not Found",
                      style: TextStyles.bodyLargeSemiBold.copyWith(
                          color: Pallete.neutral100,
                          fontSize: getFontSize(FontSizes.large)));
                } else {
                  return FutureBuilder<List<Product>>(
                    future: _productFuture,
                    builder: (context, snapshot) {
                      if (snapshot.hasData) {
                        final products = snapshot.data!;
                        return Column(
                          children: customterController.cartItems.map((item) {
                            final product = products.firstWhere(
                                (product) => product.id == item.productId);
                            return CardItemFood(
                              productId: product.id,
                              productName: product.name,
                              productImage:
                                  "http://192.168.30.215:8080/product/image/${product.image}",
                              productPrice: product.price,
                              cartItemId: item.id,
                              quantity: item.quantity,
                              onDelete: () {
                                customterController.deleteCartItem(item.id);
                                _updatePaymentSummary(products);
                              },
                              onSelectionChanged: (isSelected) {
                                setState(() {
                                  item.isChecked = isSelected;
                                  _updatePaymentSummary(products);
                                });
                              },
                              onUpdateQuantity: (newQuantity) {
                                _updateQuantity(item.id, newQuantity);
                              },
                              isChecked: item
                                  .isChecked, // Pass the isChecked value from cart item
                            );
                          }).toList(),
                        );
                      } else if (snapshot.hasError) {
                        return Text('Error: ${snapshot.error}');
                      } else {
                        return CircularProgressIndicator();
                      }
                    },
                  );
                }
              }),
              const Gap(26),
              Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  Text(
                    "Recommended For You",
                    style: TextStyles.bodyLargeSemiBold.copyWith(
                        color: Pallete.neutral100,
                        fontSize: getFontSize(FontSizes.large)),
                  ),
                  Text(
                    "See All",
                    style: TextStyles.bodyMediumMedium.copyWith(
                        color: Pallete.orangePrimary,
                        fontSize: getFontSize(FontSizes.medium)),
                  )
                ],
              ),
              const Gap(16),
              const Gap(16),
              Container(
                height: getHeight(2),
                width: double.infinity,
                color: Pallete.neutral30,
              ),
              const Gap(16),
              Container(
                width: double.infinity,
                padding: EdgeInsets.all(getSize(12)),
                decoration: BoxDecoration(
                    borderRadius: BorderRadius.circular(getSize(16)),
                    border:
                        Border.all(width: 1, color: const Color(0xFFEDEDED))),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      "Payment Summary",
                      style: TextStyles.bodyLargeSemiBold.copyWith(
                          color: Pallete.neutral100,
                          fontSize: getFontSize(FontSizes.large)),
                    ),
                    const Gap(8),
                    Row(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        Text(
                          "Total Items ($_totalItems)",
                          style: TextStyles.bodyMediumMedium.copyWith(
                              color: const Color(0xFF878787),
                              fontSize: getFontSize(FontSizes.medium)),
                        ),
                        Text(
                          "Rp ${_totalPrice.toStringAsFixed(2)}",
                          style: TextStyles.bodyMediumBold
                              .copyWith(color: Pallete.neutral100),
                        ),
                      ],
                    ),
                    const Gap(16),
                    Row(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        Text(
                          "Delivery Fee",
                          style: TextStyles.bodyMediumMedium.copyWith(
                              color: const Color(0xFF878787),
                              fontSize: getFontSize(FontSizes.medium)),
                        ),
                        Text(
                          "Free",
                          style: TextStyles.bodyMediumBold.copyWith(
                              color: Pallete.neutral100,
                              fontSize: getFontSize(FontSizes.medium)),
                        ),
                      ],
                    ),
                    const Gap(16),
                    Row(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        Text(
                          "Total",
                          style: TextStyles.bodyMediumMedium.copyWith(
                              color: const Color(0xFF878787),
                              fontSize: getFontSize(FontSizes.medium)),
                        ),
                        Text(
                          "Rp ${_totalPrice.toStringAsFixed(2)}",
                          style: TextStyles.bodyMediumBold.copyWith(
                              color: Pallete.neutral100,
                              fontSize: getFontSize(FontSizes.medium)),
                        ),
                      ],
                    ),
                  ],
                ),
              ),
              const Gap(26),
              DefaultButton(btnContent: "Order Now"),
              const Gap(6)
            ],
          ),
        ),
      ),
    );
  }
}
