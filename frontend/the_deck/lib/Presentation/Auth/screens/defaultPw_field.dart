import 'package:flutter/material.dart';
import 'package:gap/gap.dart';
import 'package:the_deck/Core/app_colors.dart';
import 'package:the_deck/Core/response_conf.dart';
import 'package:the_deck/Core/text_styles.dart';

class DefaultFieldPW extends StatefulWidget {
  DefaultFieldPW({
    Key? key,
    required this.hintText,
    this.labelText,
    this.prefixIcon,
    this.suffixIcon,
    this.isPasswordField,
    this.controller,
  }) : super(key: key);

  final String hintText;
  final String? labelText;
  final bool? isPasswordField;
  final IconData? suffixIcon;
  final IconData? prefixIcon;
  final TextEditingController? controller;

  @override
  State<DefaultFieldPW> createState() => _DefaultFieldPWState();
}

class _DefaultFieldPWState extends State<DefaultFieldPW> {
  bool isHideCaracter = true;

  @override
  void initState() {
    super.initState();
    isHideCaracter = widget.isPasswordField ?? false;
  }

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        widget.labelText != null
            ? Text(
                "${widget.labelText}",
                style: TextStyles.bodyMediumMedium.copyWith(
                    color: Pallete.neutral100, fontSize: getFontSize(14)),
              )
            : const SizedBox(),
        const Gap(8),
        TextFormField(
          controller: widget.controller,
          obscureText: isHideCaracter,
          obscuringCharacter: "*",
          style: TextStyles.bodyMediumMedium
              .copyWith(color: Pallete.neutral100, fontSize: getFontSize(14)),
          decoration: InputDecoration(
              hintText: widget.hintText,
              hintStyle: TextStyles.bodyMediumMedium.copyWith(
                  color: Pallete.neutral60, fontSize: getFontSize(14)),
              contentPadding: EdgeInsets.all(getSize(16)),
              border: OutlineInputBorder(
                borderSide:
                    const BorderSide(width: 1, color: Pallete.neutral40),
                borderRadius: BorderRadius.circular(getSize(8)),
              ),
              enabledBorder: OutlineInputBorder(
                  borderSide:
                      const BorderSide(width: 1, color: Pallete.neutral40),
                  borderRadius: BorderRadius.circular(getSize(8))),
              focusedBorder: OutlineInputBorder(
                  borderSide:
                      const BorderSide(width: 1, color: Pallete.neutral40),
                  borderRadius: BorderRadius.circular(getSize(8))),
              prefixIcon: widget.prefixIcon != null
                  ? Icon(
                      widget.prefixIcon,
                      size: getSize(20),
                      color: const Color(0xFF878787),
                    )
                  : null,
              suffixIcon: widget.isPasswordField == true
                  ? InkWell(
                      onTap: () => setState(() {
                        isHideCaracter = !isHideCaracter;
                      }),
                      child: Icon(
                        isHideCaracter ? Icons.visibility_off : Icons.visibility,
                        size: getSize(20),
                        color: Pallete.neutral100,
                      ),
                    )
                  : widget.suffixIcon != null
                      ? Icon(
                          widget.suffixIcon,
                          size: getSize(20),
                          color: const Color(0xFF878787),
                        )
                      : null),
        ),
      ],
    );
  }
}
