import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:the_deck/Controller/CustomerController.dart';
import 'package:the_deck/Core/app_colors.dart';
import 'package:the_deck/Core/assets_constantes.dart';
import 'package:the_deck/Core/font_size.dart';
import 'package:the_deck/Core/response_conf.dart';
import 'package:the_deck/Core/text_styles.dart';
import 'package:gap/gap.dart';
import 'package:the_deck/Presentation/Foods/Views/about_menu_view.dart';

class CardItemFood extends StatefulWidget {
  final int cartItemId;
  final int productId;
  final String productName;
  final String productImage;
  final double productPrice;
  int quantity;
  final VoidCallback onDelete;
  final ValueChanged<bool> onSelectionChanged;
  final ValueChanged<int> onUpdateQuantity;
  final bool isChecked;

  CardItemFood({
    Key? key,
    required this.cartItemId,
    required this.productId,
    required this.productName,
    required this.productImage,
    required this.productPrice,
    required this.onDelete,
    required this.quantity,
    required this.onSelectionChanged,
    required this.onUpdateQuantity,
    required this.isChecked,
  }) : super(key: key);

  @override
  _CardItemFoodState createState() => _CardItemFoodState();
}

class _CardItemFoodState extends State<CardItemFood> {
  bool isChecked = false;
  final RegisterController registerController = Get.find<RegisterController>();

  void _updateQuantity(int newQuantity) async {
    await registerController.updateCartItemQuantity(
        widget.cartItemId, newQuantity);
    setState(() {
      widget.quantity = newQuantity;
    });
    widget.onSelectionChanged(isChecked);
  }

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: () {
        setState(() {
          isChecked = !isChecked;
        });
        widget.onSelectionChanged(isChecked);
      },
      child: Padding(
        padding: EdgeInsets.only(top: getHeight(16)),
        child: Container(
          width: double.infinity,
          height: getHeight(108),
          padding: EdgeInsets.all(getSize(12)),
          decoration: ShapeDecoration(
            color: Colors.white,
            shape: RoundedRectangleBorder(
              borderRadius: BorderRadius.circular(getSize(12)),
            ),
            shadows: const [
              BoxShadow(
                color: Color(0x0A000000),
                blurRadius: 60,
                offset: Offset(6, 6),
                spreadRadius: 0,
              )
            ],
          ),
          child: Row(
            children: [
              Checkbox(
                shape: RoundedRectangleBorder(
                  borderRadius: BorderRadius.circular(getSize(4)),
                ),
                value: isChecked,
                onChanged: (value) {
                  setState(() {
                    isChecked = value!;
                  });
                  widget.onSelectionChanged(isChecked);
                },
                fillColor: MaterialStateProperty.all(
                    isChecked ? Pallete.orangePrimary : Colors.transparent),
              ),
              const Gap(16),
              GestureDetector(
                onTap: () {
                  Get.to(() => AboutMenuView(productId: widget.productId));
                },
                child: Container(
                  width: getWidth(85),
                  height: getHeight(82),
                  decoration: BoxDecoration(
                    borderRadius: BorderRadius.circular(getSize(8)),
                    image: DecorationImage(
                      image: NetworkImage(widget.productImage),
                      fit: BoxFit.fill,
                    ),
                  ),
                ),
              ),
              const Gap(16),
              Expanded(
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(widget.productName,
                        textAlign: TextAlign.center,
                        style: TextStyles.bodyLargeSemiBold.copyWith(
                            color: Pallete.neutral100,
                            fontSize: getFontSize(FontSizes.large))),
                    const Gap(4),
                    Text('Rp ${widget.productPrice.toStringAsFixed(2)}',
                        textAlign: TextAlign.center,
                        style: TextStyles.bodyMediumBold.copyWith(
                            color: Pallete.orangePrimary,
                            fontSize: getFontSize(FontSizes.medium))),
                    const Gap(8),
                    Row(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        Row(
                          children: [
                            GestureDetector(
                              onTap: () {
                                if (widget.quantity > 1) {
                                  widget.onUpdateQuantity(widget.quantity - 1);
                                }
                              },
                              child: Container(
                                height: getSize(28),
                                width: getSize(28),
                                decoration: BoxDecoration(
                                    shape: BoxShape.circle,
                                    border: Border.all(
                                        color: const Color(0xFFEAEAEA),
                                        width: 1)),
                                child: Icon(
                                  Icons.remove,
                                  size: getSize(24),
                                  weight: 2,
                                  color: const Color(0xFF878787),
                                ),
                              ),
                            ),
                            const Gap(8),
                            Text('${widget.quantity}',
                                style: TextStyles.bodyLargeBold.copyWith(
                                    color: Pallete.neutral100,
                                    fontSize: getFontSize(FontSizes.large))),
                            const Gap(16),
                            GestureDetector(
                              onTap: () {
                                _updateQuantity(widget.quantity + 1);
                              },
                              child: Container(
                                height: getSize(28),
                                width: getSize(28),
                                decoration: BoxDecoration(
                                    shape: BoxShape.circle,
                                    border: Border.all(
                                        color: const Color(0xFFEAEAEA),
                                        width: 1)),
                                child: Icon(
                                  Icons.add,
                                  size: getSize(24),
                                  weight: 2,
                                  color: Pallete.neutral100,
                                ),
                              ),
                            ),
                          ],
                        ),
                        GestureDetector(
                          onTap: () {
                            widget.onDelete();
                          },
                          child: Icon(
                            Icons.delete,
                            color: Pallete.pureError,
                            size: getSize(20),
                          ),
                        ),
                      ],
                    )
                  ],
                ),
              )
            ],
          ),
        ),
      ),
    );
  }
}
