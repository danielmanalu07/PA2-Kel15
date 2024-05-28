import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:the_deck/Controller/CategoryController.dart';
import 'package:the_deck/Core/app_colors.dart';
import 'package:the_deck/Core/font_size.dart';
import 'package:the_deck/Core/response_conf.dart';
import 'package:the_deck/Core/text_styles.dart';
import 'package:the_deck/Presentation/Product/views/Product_list.dart';
import 'package:gap/gap.dart';

class CategoryListScreen extends StatelessWidget {
  final CategoryController categoryController = Get.put(CategoryController());

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('All Categories'),
      ),
      body: Obx(() {
        if (categoryController.isLoading.value) {
          return Center(child: CircularProgressIndicator());
        }
        return ListView.builder(
          itemCount: categoryController.categories.length,
          itemBuilder: (context, index) {
            final category = categoryController.categories[index];
            return GestureDetector(
              onTap: () => Get.to(() => ProductListScreen(
                    categoryId: category.id,
                  )),
              child: Container(
                margin: EdgeInsets.all(8),
                padding: EdgeInsets.all(16),
                decoration: BoxDecoration(
                  gradient: LinearGradient(
                    colors: [
                      Colors.orange,
                      Colors.deepOrange,
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
                  borderRadius: BorderRadius.circular(16),
                ),
                child: Row(
                  children: [
                    Icon(
                      Icons.fastfood, // Replace with actual category icon if available
                      color: Colors.white,
                      size: getSize(40),
                    ),
                    const Gap(16),
                    Text(
                      category.name,
                      style: TextStyle(
                        color: Colors.white,
                        fontWeight: FontWeight.bold,
                        fontSize: getFontSize(FontSizes.medium),
                      ),
                    ),
                  ],
                ),
              ),
            );
          },
        );
      }),
    );
  }
}
