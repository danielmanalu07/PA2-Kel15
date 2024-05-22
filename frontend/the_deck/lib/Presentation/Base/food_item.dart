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
              child: Align(
                alignment: Alignment.topRight,
                child: Container(
                  width: getSize(24),
                  height: getSize(24),
                  margin: EdgeInsets.all(getSize(4)),
                  decoration: BoxDecoration(
                    shape: BoxShape.circle,
                    color: Colors.white,
                  ),
                  child: Icon(
                    Icons.favorite_border,
                    color: Pallete.pureError,
                    size: getSize(20),
                  ),
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
            const Gap(4),
            Row(
              children: [
                Row(
                  children: [
                    Icon(
                      Icons.star,
                      color: Pallete.orangePrimary,
                      size: getSize(16),
                    ),
                    const Gap(4),
                    Text(
                      "4.9",
                      style: TextStyles.bodySmallMedium.copyWith(
                        color: Pallete.neutral100,
                      ),
                    ),
                  ],
                ),
                const Spacer(),
                Row(
                  children: [
                    Icon(
                      Icons.location_on_outlined,
                      color: Pallete.orangePrimary,
                      size: getSize(16),
                    ),
                    const Gap(4),
                    Text(
                      "190m",
                      style: TextStyles.bodySmallMedium.copyWith(
                        color: Pallete.neutral100,
                        fontSize: getFontSize(FontSizes.small),
                      ),
                    ),
                  ],
                ),
              ],
            ),
            const Gap(6),
            Text(
              'Rp ${price.toStringAsFixed(2)}', // Menampilkan harga dengan dua desimal
              style: TextStyles.bodyLargeBold.copyWith(
                color: Pallete.orangePrimary,
                fontSize: getFontSize(FontSizes.large),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
