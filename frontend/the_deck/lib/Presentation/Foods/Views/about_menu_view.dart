import 'package:get/get.dart';
import 'package:the_deck/Controller/CustomerController.dart';
import 'package:the_deck/Controller/ProductController.dart';
import 'package:the_deck/Core/app_colors.dart';
import 'package:the_deck/Core/assets_constantes.dart';
import 'package:the_deck/Core/font_size.dart';
import 'package:the_deck/Core/response_conf.dart';
import 'package:the_deck/Core/text_styles.dart';
import 'package:the_deck/Presentation/Auth/screens/default_button.dart';
import 'package:the_deck/Presentation/Base/base.dart';
import 'package:flutter/material.dart';
import 'package:gap/gap.dart';

class AboutMenuView extends StatefulWidget {
  const AboutMenuView({Key? key, required this.productId}) : super(key: key);

  final int productId;

  @override
  State<AboutMenuView> createState() => _AboutMenuViewState();
}

class _AboutMenuViewState extends State<AboutMenuView> {
  final ProductController productController = Get.put(ProductController());
  final RegisterController customerController = Get.put(RegisterController());
  var quantity = 1.obs;

  @override
  void initState() {
    super.initState();
    productController.getProductById(widget.productId);
  }

  void addToCart() async {
    await customerController.addToCart(widget.productId, quantity.value);
  }

  @override
  Widget build(BuildContext context) {
    MathUtils.init(context);
    return Scaffold(
      appBar:
          buildAppBar(buildContext: context, screenTitle: "About This Menu"),
      body: Padding(
        padding: EdgeInsets.symmetric(horizontal: getWidth(24)),
        child: Obx(
          () {
            if (productController.product.value == null) {
              return Center(
                child: CircularProgressIndicator(),
              );
            } else {
              final product = productController.product.value!;
              return SingleChildScrollView(
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    const Gap(8),
                    Container(
                      width: double.infinity,
                      height: getHeight(295),
                      padding: EdgeInsets.all(getSize(16)),
                      decoration: ShapeDecoration(
                        image: DecorationImage(
                          image: NetworkImage(
                              "http://172.26.43.150:8080/product/image/${product.image}"),
                          fit: BoxFit.fill,
                        ),
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(16),
                        ),
                      ),
                    ),
                    const Gap(16),
                    Text(
                      product.name,
                      style: TextStyles.headingH5SemiBold.copyWith(
                        color: Pallete.neutral100,
                        fontSize: getFontSize(FontSizes.h5),
                      ),
                    ),
                    const Gap(8),
                    Text(
                      "Rp ${product.price.toStringAsFixed(2)}",
                      style: TextStyles.headingH6Bold.copyWith(
                        color: Pallete.greenStrong,
                        fontSize: getFontSize(FontSizes.h6),
                      ),
                    ),
                    const Gap(16),
                    Container(
                      width: double.infinity,
                      height: 2,
                      color: const Color(0xFFEDEDED),
                    ),
                    const Gap(16),
                    Text(
                      "Description",
                      style: TextStyles.headingH5SemiBold.copyWith(
                        color: Pallete.neutral100,
                        fontSize: getFontSize(FontSizes.h5),
                      ),
                    ),
                    const Gap(8),
                    Text(
                      product.description,
                      style: TextStyles.bodyMediumRegular.copyWith(
                        color: const Color(0xFF878787),
                        fontSize: getFontSize(FontSizes.medium),
                      ),
                    ),
                    const Gap(18),
                  ],
                ),
              );
            }
          },
        ),
      ),
      bottomSheet: Container(
        padding: EdgeInsets.symmetric(horizontal: getWidth(24))
            .copyWith(top: getHeight(16), bottom: getHeight(32)),
        height: getHeight(100),
        width: double.infinity,
        decoration: const BoxDecoration(
          color: Colors.white,
        ),
        child: Row(
          children: [
            const Gap(26),
            Expanded(
              child: DefaultButton(
                btnContent: "Add to Cart",
                btnIcon: Icons.shopping_cart_outlined,
                function: addToCart,
              ),
            ),
          ],
        ),
      ),
    );
  }
}
