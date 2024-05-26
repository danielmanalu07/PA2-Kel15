import 'package:get/get.dart';
import 'package:get/get_connect/http/src/utils/utils.dart';
import 'package:the_deck/Controller/CustomerController.dart';
import 'package:the_deck/Core/Routes/routes_name.dart';
import 'package:the_deck/Core/app_colors.dart';
import 'package:the_deck/Core/response_conf.dart';
import 'package:the_deck/Core/text_styles.dart';
import 'package:the_deck/Presentation/Auth/screens/defaultPw_field.dart';
import 'package:the_deck/Presentation/Auth/screens/default_button.dart';
import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';
import '../screens/default_field.dart';
import 'package:gap/gap.dart';

class LoginView extends StatelessWidget {
  const LoginView({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    final TextEditingController _emailController = TextEditingController();
    final TextEditingController _passwordController = TextEditingController();
    final RegisterController _customerControlller =
        Get.put(RegisterController());

    void _login() async {
      try {
        await _customerControlller.loginUser(
            _emailController.text, _passwordController.text);
      } catch (error) {
        Get.snackbar('Error', 'Email and password are incorrect');
      }
    }

    MathUtils.init(context);
    return Scaffold(
        body: Padding(
      padding: EdgeInsets.symmetric(horizontal: getWidth(24)).copyWith(
        top: MediaQuery.of(context).viewPadding.top,
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          const Gap(32),
          Text(
            "Login to your\n account.",
            style: TextStyles.headingH4SemiBold
                .copyWith(color: Pallete.neutral100),
          ),
          const Gap(8),
          Text(
            "Please sign in to your account ",
            style: TextStyles.bodyMediumMedium
                .copyWith(color: Pallete.neutral60, fontSize: getFontSize(14)),
          ),
          const Gap(32),
          DefaultField(
            hintText: "Enter Email",
            controller: _emailController,
            labelText: "Email",
          ),
          const Gap(14),
          DefaultFieldPW(
            hintText: "Password",
            labelText: "Password",
            controller: _passwordController,
            isPasswordField: true,
          ),
          const Gap(24),
          Align(
            alignment: Alignment.topRight,
            child: InkWell(
              onTap: () =>
                  Navigator.pushNamed(context, RoutesName.forgetPassword),
              child: Text(
                "Forgot password?",
                style: TextStyles.bodyMediumMedium.copyWith(
                    color: Pallete.orangePrimary, fontSize: getFontSize(14)),
              ),
            ),
          ),
          const Gap(24),
          DefaultButton(
            btnContent: "Sign In",
            function: _login,
          ),
          const Gap(32),
          Align(
            alignment: Alignment.center,
            child: Text.rich(
              TextSpan(
                children: [
                  TextSpan(
                    text: "Don't have an account?",
                    style: TextStyles.bodyMediumMedium.copyWith(
                        color: Pallete.neutral100, fontSize: getFontSize(14)),
                  ),
                  TextSpan(
                      text: ' ',
                      style: TextStyles.bodyMediumSemiBold
                          .copyWith(fontSize: getFontSize(14))),
                  TextSpan(
                      recognizer: TapGestureRecognizer()
                        ..onTap = () => Navigator.pushReplacementNamed(
                            context, RoutesName.signUp),
                      text: 'Register',
                      style: TextStyles.bodyMediumSemiBold.copyWith(
                          color: Pallete.orangePrimary,
                          fontSize: getFontSize(14))),
                ],
              ),
            ),
          )
        ],
      ),
    ));
  }
}
