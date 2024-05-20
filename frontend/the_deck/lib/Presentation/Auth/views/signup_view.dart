import 'dart:io';

import 'package:the_deck/Core/Routes/routes_name.dart';
import 'package:the_deck/Core/Utils/utils.dart';
import 'package:the_deck/Core/app_colors.dart';
import 'package:the_deck/Core/font_size.dart';
import 'package:the_deck/Core/response_conf.dart';
import 'package:the_deck/Core/text_styles.dart';
import 'package:the_deck/Presentation/Auth/screens/account_status.dart';
import 'package:the_deck/Presentation/Auth/screens/default_button.dart';
import 'package:the_deck/Presentation/Auth/screens/default_field.dart';
import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:gap/gap.dart';
import 'package:image_picker/image_picker.dart';

class SignUpView extends StatefulWidget {
  const SignUpView({Key? key}) : super(key: key);

  @override
  _SignUpViewState createState() => _SignUpViewState();
}

class _SignUpViewState extends State<SignUpView> {
  final emailController = TextEditingController();
  final passwordController = TextEditingController();
  final usernameController = TextEditingController();
  final fullNameController = TextEditingController();
  final phoneController = TextEditingController();
  final addressController = TextEditingController();
  final genderController = TextEditingController();
  final dateOfBirthController = TextEditingController();
  final ImagePicker _picker = ImagePicker();
  XFile? _image;

  @override
  Widget build(BuildContext context) {
    MathUtils.init(context);

    return Scaffold(
      resizeToAvoidBottomInset: false,
      body: Padding(
        padding: EdgeInsets.symmetric(horizontal: getWidth(24)).copyWith(
          top: MediaQuery.of(context).viewPadding.top,
        ),
        child: SingleChildScrollView(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              const Gap(32),
              Text(
                "Create your new \naccount",
                style: TextStyles.headingH4SemiBold.copyWith(
                  color: Pallete.neutral100,
                  fontSize: getFontSize(FontSizes.h4),
                ),
              ),
              const Gap(8),
              Text(
                "Create an account to start looking for the food \nyou like",
                style: TextStyles.bodyMediumMedium.copyWith(
                  color: Pallete.neutral60,
                  fontSize: getFontSize(FontSizes.medium),
                ),
              ),
              const Gap(12),
              DefaultField(
                hintText: "Full Name",
                controller: fullNameController,
                labelText: "Full Name",
              ),
              const Gap(14),
              DefaultField(
                hintText: "Enter Email",
                controller: emailController,
                labelText: "Email Address",
              ),
              const Gap(14),
              DefaultField(
                hintText: "User Name",
                controller: usernameController,
                labelText: "User Name",
              ),
              const Gap(14),
              DefaultField(
                hintText: "Password",
                controller: passwordController,
                labelText: "Password",
                isPasswordField: true,
              ),
              const Gap(14),
              DefaultField(
                hintText: "Phone Number",
                controller: phoneController,
                labelText: "Phone Number",
              ),
              const Gap(14),
              DefaultField(
                hintText: "Address",
                controller: addressController,
                labelText: "Address",
              ),
              const Gap(14),
              DefaultField(
                hintText: "Gender",
                controller: genderController,
                labelText: "Gender",
              ),
              const Gap(14),
              GestureDetector(
                onTap: () async {
                  DateTime? pickedDate = await showDatePicker(
                    context: context,
                    initialDate: DateTime.now(),
                    firstDate: DateTime(1900),
                    lastDate: DateTime(2100),
                  );
                  if (pickedDate != null) {
                    setState(() {
                      dateOfBirthController.text =
                          "${pickedDate.year}-${pickedDate.month.toString().padLeft(2, '0')}-${pickedDate.day.toString().padLeft(2, '0')}";
                    });
                  }
                },
                child: AbsorbPointer(
                  child: DefaultField(
                    hintText: "Date of Birth (YYYY-MM-DD)",
                    controller: dateOfBirthController,
                    labelText: "Date of Birth",
                  ),
                ),
              ),
              const Gap(14),
              GestureDetector(
                onTap: () async {
                  final XFile? pickedFile = await _picker.pickImage(
                    source: ImageSource.gallery,
                  );
                  setState(() {
                    _image = pickedFile;
                  });
                },
                child: Container(
                  width: double.infinity,
                  height: getHeight(150),
                  decoration: BoxDecoration(
                    color: Pallete.neutral20,
                    borderRadius: BorderRadius.circular(8),
                    border: Border.all(color: Pallete.neutral40),
                  ),
                  child: _image != null
                      ? Image.file(
                          File(_image!.path),
                          fit: BoxFit.cover,
                        )
                      : Column(
                          mainAxisAlignment: MainAxisAlignment.center,
                          children: [
                            Icon(
                              Icons.camera_alt,
                              color: Pallete.neutral60,
                              size: getSize(50),
                            ),
                            const Gap(8),
                            Text(
                              "Upload Profile Picture",
                              style: TextStyles.bodyMediumMedium.copyWith(
                                color: Pallete.neutral60,
                                fontSize: getFontSize(FontSizes.medium),
                              ),
                            ),
                          ],
                        ),
                ),
              ),
              const Gap(24),
              const Gap(24),
              DefaultButton(
                btnContent: "Register",
              ),
              const Gap(24),
              Align(
                alignment: Alignment.center,
                child: Text.rich(
                  TextSpan(
                    children: [
                      TextSpan(
                        recognizer: TapGestureRecognizer()
                          ..onTap = () => Navigator.pushReplacementNamed(
                              context, RoutesName.login),
                        text: 'Sign In',
                        style: TextStyles.bodyMediumSemiBold.copyWith(
                          color: Pallete.orangePrimary,
                          fontSize: getFontSize(14),
                        ),
                      ),
                    ],
                  ),
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
