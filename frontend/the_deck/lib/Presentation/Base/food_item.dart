import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:the_deck/Core/Routes/routes_name.dart';
import 'package:the_deck/Core/app_colors.dart';
import 'package:the_deck/Core/font_size.dart';
import 'package:the_deck/Core/response_conf.dart';
import 'package:the_deck/Core/text_styles.dart';
import 'package:gap/gap.dart';
import 'package:the_deck/Presentation/Foods/Views/about_menu_view.dart';

class FoodItem extends StatelessWidget {
  const FoodItem({
    Key? key,
    required this.id,
    required this.name,
    required this.image,
    required this.price,
  }) : super(key: key);

  final String name;
  final String image;
  final double price;
  final int id;

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: () {
        Get.to(() => AboutMenuView(productId: id));
      },
      child: Container(
        height: getHeight(204),
        width: getWidth(153),
        padding: EdgeInsets.all(getSize(8)),
        decoration: BoxDecoration(
          borderRadius: BorderRadius.circular(getSize(12)),
          boxShadow: const [
            BoxShadow(
              color: Color(0x0A000000),
              blurRadius: 12, // Mengurangi blur radius
              offset: Offset(2, 2), // Mengurangi offset
              spreadRadius: 0,
            )
          ],
          color: Colors.white,
        ),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Container(
              width: double.infinity,
              height: getHeight(106),
              decoration: BoxDecoration(
                borderRadius: BorderRadius.circular(getSize(8)),
                image: DecorationImage(
                  image: NetworkImage(image),
                  fit: BoxFit.cover,
                ),
              ),
            ),
            const Gap(8),
            Text(
              name,
              style: TextStyles.bodyLargeMedium.copyWith(
                color: Pallete.neutral100,
                fontSize: getFontSize(FontSizes.large),
              ),
            ),
        
            const Gap(6),
            Text(
              'Rp ${price.toStringAsFixed(2)}', // Menampilkan harga dengan dua desimal
              style: TextStyles.bodyLargeBold.copyWith(
                color: Pallete.greenStrong,
                fontSize: getFontSize(FontSizes.large),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
