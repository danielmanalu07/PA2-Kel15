import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:the_deck/Controller/CategoryController.dart';
import 'package:the_deck/Controller/CustomerController.dart';
import 'package:the_deck/Controller/ProductController.dart';
import 'package:the_deck/Core/Routes/routes_name.dart';
import 'package:the_deck/Core/app_colors.dart';
import 'package:the_deck/Core/assets_constantes.dart';
import 'package:the_deck/Core/font_size.dart';
import 'package:the_deck/Core/response_conf.dart';
import 'package:the_deck/Core/text_styles.dart';
import 'package:the_deck/Models/Product.dart';
import 'package:the_deck/Presentation/Base/food_item.dart';
import 'package:gap/gap.dart';
import 'package:the_deck/Presentation/Category/CategoryList.dart';
import 'package:the_deck/Presentation/Models/category_model.dart';
import 'package:the_deck/Presentation/Product/views/Product_list.dart';

class HomeView extends StatefulWidget {
  HomeView({Key? key}) : super(key: key);

  @override
  State<HomeView> createState() => _HomeViewState();
}

class _HomeViewState extends State<HomeView> {
  final CategoryController categoryController = Get.put(CategoryController());
  final ProductController _productController = ProductController();
  final RegisterController _controller = Get.put(RegisterController());

  late Future<List<Product>> _productFuture;

  @override
  void initState() {
    super.initState();
    _productFuture = _productController.getProductList();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Obx(() {
        final customer = _controller.userProfile.value;
        if (categoryController.isLoading.value) {
          return Center(child: CircularProgressIndicator());
        }
        return SingleChildScrollView(
          child: Column(
            children: [
              Container(
                height: getHeight(250),
                width: double.infinity,
                padding: EdgeInsets.symmetric(
                  horizontal: getWidth(24),
                ).copyWith(
                  top: MediaQuery.of(context).viewPadding.top,
                ),
                decoration: const BoxDecoration(
                    image: DecorationImage(
                        image:
                            AssetImage(AssetsConstants.homeTopBackgroundImage),
                        fit: BoxFit.fill)),
                child: Padding(
                  padding: EdgeInsets.only(
                      top: getHeight(20), bottom: getHeight(20)),
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Row(
                        mainAxisAlignment: MainAxisAlignment.spaceBetween,
                        children: [
                          Column(
                            children: [
                              Row(
                                children: [
                                  // Replace "Your Location" with customer's name
                                  Text(
                                    customer?.name ?? 'Guest',
                                    style: TextStyles.bodyMediumRegular
                                        .copyWith(
                                            color: Colors.white,
                                            fontSize:
                                                getFontSize(FontSizes.medium)),
                                  ),
                                ],
                              ),
                            ],
                          ),
                        ],
                      ),
                      const Gap(26),
                      Text(
                        "The Deck Provide the best \nfood for you",
                        style: TextStyles.headingH4SemiBold.copyWith(
                            color: Pallete.neutral10,
                            fontSize: getFontSize(FontSizes.h4)),
                      ),
                    ],
                  ),
                ),
              ),
              Padding(
                padding: EdgeInsets.symmetric(horizontal: getWidth(24)),
                child: Column(
                  children: [
                    const Gap(26),
                    Row(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        Text(
                          "Find by Category",
                          style: TextStyles.bodyLargeSemiBold.copyWith(
                              color: Pallete.neutral100,
                              fontSize: getFontSize(FontSizes.large)),
                        ),
                        GestureDetector(
                          onTap: () => Get.to(() =>
                              CategoryListScreen()), // Navigate to CategoryListScreen
                          child: Text(
                            "See All",
                            style: TextStyles.bodyMediumMedium.copyWith(
                                color: Pallete.greenStrong,
                                fontSize: getFontSize(FontSizes.medium)),
                          ),
                        ),
                      ],
                    ),
                    const Gap(18),
                    SingleChildScrollView(
                      scrollDirection: Axis.horizontal,
                      child: Row(
                        children: categoryController.categories
                            .take(4)
                            .map(
                              (category) => GestureDetector(
                                onTap: () => Get.to(() => ProductListScreen(
                                      categoryId: category.id,
                                    )),
                                child: Container(
                                  margin: EdgeInsets.only(right: getSize(8)),
                                  width: getSize(80), // Increased width
                                  height: getSize(100), // Increased height
                                  padding: const EdgeInsets.all(8),
                                  decoration: BoxDecoration(
                                    gradient: LinearGradient(
                                      colors: [
                                        Colors.green,
                                        Colors.greenAccent,
                                      ],
                                      begin: Alignment.topLeft,
                                      end: Alignment.bottomRight,
                                    ),
                                    boxShadow: const [
                                      BoxShadow(
                                        color: Color(0x0A111111),
                                        blurRadius: 24,
                                        offset: Offset(0, 4),
                                        spreadRadius: 0,
                                      )
                                    ],
                                    borderRadius: BorderRadius.circular(
                                        16), // Increased border radius
                                  ),
                                  child: Column(
                                    mainAxisAlignment: MainAxisAlignment.center,
                                    children: [
                                      Icon(
                                        Icons
                                            .fastfood, // Replace with actual category icon if available
                                        color: Colors.white,
                                        size: getSize(40),
                                      ),
                                      const Gap(8),
                                      Text(
                                        category.name,
                                        style: TextStyle(
                                          color: Colors.white,
                                          fontWeight: FontWeight.bold,
                                          fontSize:
                                              getFontSize(FontSizes.medium),
                                        ),
                                        textAlign: TextAlign.center,
                                      ),
                                    ],
                                  ),
                                ),
                              ),
                            )
                            .toList(),
                      ),
                    ),
                    const Gap(24),
                    FutureBuilder<List<Product>>(
                      future: _productFuture,
                      builder: (context, snapshot) {
                        if (snapshot.connectionState ==
                            ConnectionState.waiting) {
                          return Center(child: CircularProgressIndicator());
                        } else if (snapshot.hasError) {
                          return Center(
                              child: Text('Error: ${snapshot.error}'));
                        } else {
                          return GridView.builder(
                            shrinkWrap: true,
                            physics: NeverScrollableScrollPhysics(),
                            gridDelegate:
                                SliverGridDelegateWithFixedCrossAxisCount(
                              crossAxisCount: 2,
                              crossAxisSpacing: getSize(16),
                              mainAxisSpacing: getSize(16),
                              childAspectRatio: 3 / 4,
                            ),
                            itemCount: snapshot.data!.length,
                            itemBuilder: (context, index) {
                              final product = snapshot.data![index];
                              return FoodItem(
                                id: product.id,
                                name: product.name,
                                image:
                                    "http://192.168.30.215:8080/product/image/${product.image}",
                                price: product.price,
                              );
                            },
                          );
                        }
                      },
                    ),
                  ],
                ),
              )
            ],
          ),
        );
      }),
    );
  }
}
