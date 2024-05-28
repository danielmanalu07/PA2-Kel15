import 'package:flutter/gestures.dart';
import 'package:get/get.dart';
import 'package:the_deck/Controller/CustomerController.dart';
import 'package:the_deck/Core/Routes/routes_name.dart';
import 'package:the_deck/Core/app_colors.dart';
import 'package:the_deck/Core/assets_constantes.dart';
import 'package:the_deck/Core/font_size.dart';
import 'package:the_deck/Core/response_conf.dart';
import 'package:the_deck/Core/text_styles.dart';
import 'package:the_deck/Presentation/Base/base.dart';
import 'package:the_deck/Presentation/Cart/MyOrder.dart';
import 'package:the_deck/Presentation/Profil/screens/profile_info_tile.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:gap/gap.dart';

class ProfilView extends StatefulWidget {
  const ProfilView({Key? key}) : super(key: key);

  @override
  State<ProfilView> createState() => _ProfilViewState();
}

class _ProfilViewState extends State<ProfilView> {
  final RegisterController _controller = Get.put(RegisterController());

  @override
  void initState() {
    super.initState();
    _controller.getUserProfile();
  }

  @override
  Widget build(BuildContext context) {
    return Obx(() {
      final customer = _controller.userProfile.value;
      if (customer == null) {
        return Scaffold(
          appBar: buildAppBar(
              buildContext: context,
              screenTitle: "Profile Settings",
              isBackup: false),
          body: Center(child: CircularProgressIndicator()),
        );
      }
      return Scaffold(
        appBar: buildAppBar(
            buildContext: context,
            screenTitle: "Profile Settings",
            isBackup: false),
        body: Padding(
          padding: EdgeInsets.symmetric(horizontal: getWidth(24)),
          child: SingleChildScrollView(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.center,
              children: [
                const Gap(24),
                Stack(
                  children: [
                    CircleAvatar(
                      backgroundImage: NetworkImage(
                          'http://192.168.30.215:8080/customer/image/${customer.image}'),
                      radius: getSize(50),
                    ),
                    Positioned(
                      left: getSize(72),
                      bottom: getSize(8),
                      child: Container(
                        width: getSize(32),
                        height: getSize(32),
                        padding: EdgeInsets.all(getSize(6)),
                        decoration: const BoxDecoration(
                            color: Color(0xFFF5F5FF), shape: BoxShape.circle),
                        child: Icon(
                          CupertinoIcons.camera_fill,
                          color: Pallete.orangePrimary,
                          size: getSize(20),
                        ),
                      ),
                    )
                  ],
                ),
                const Gap(16),
                Text(
                  customer.name,
                  style: TextStyles.bodyLargeSemiBold.copyWith(
                      color: Pallete.neutral100,
                      fontSize: getFontSize(FontSizes.large)),
                ),
                const Gap(4),
                Text(
                  customer.email,
                  style: TextStyles.bodyMediumRegular.copyWith(
                      color: const Color(0xFF878787),
                      fontSize: getFontSize(FontSizes.medium)),
                ),
                const Gap(24),
                Container(
                  width: double.infinity,
                  decoration: const ShapeDecoration(
                    shape: RoundedRectangleBorder(
                      side: BorderSide(
                        width: 1,
                        strokeAlign: BorderSide.strokeAlignCenter,
                        color: Color(0xFFEDEDED),
                      ),
                    ),
                  ),
                ),
                const Gap(24),
                Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      "Profile",
                      style: TextStyles.bodySmallMedium.copyWith(
                          color: const Color(0xFF878787),
                          fontSize: getFontSize(FontSizes.small)),
                    ),
                    ProfileInfoTile(
                        function: () => Navigator.pushNamed(
                            context, RoutesName.personnalData),
                        prefixIcon: Icons.person,
                        title: "Personal Data"),
                    ProfileInfoTile(
                        function: () => Get.to(() => MyOrder()),
                        prefixIcon: Icons.shopping_bag,
                        title: "My Order"),
                    ProfileInfoTile(
                        function: () =>
                            Navigator.pushNamed(context, RoutesName.settings),
                        prefixIcon: Icons.settings,
                        title: "Settings"),
                    ProfileInfoTile(
                        function: () =>
                            Navigator.pushNamed(context, RoutesName.extraCard),
                        prefixIcon: Icons.credit_card,
                        title: "Extra Card"),
                  ],
                ),
                const Gap(16),
                Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      "Support",
                      style: TextStyles.bodySmallMedium.copyWith(
                          color: const Color(0xFF878787),
                          fontSize: getFontSize(FontSizes.small)),
                    ),
                    ProfileInfoTile(
                        prefixIcon: Icons.info_outline, title: "Help Center"),
                    ProfileInfoTile(
                        prefixIcon: Icons.delete,
                        title: "Request Account Deletion"),
                    ProfileInfoTile(
                        prefixIcon: Icons.person_add,
                        title: "Add another account"),
                  ],
                ),
                const Gap(16),
                Align(
                  alignment: Alignment.center,
                  child: Text.rich(
                    TextSpan(
                      children: [
                        TextSpan(
                          recognizer: TapGestureRecognizer()
                            ..onTap = _controller.logout,
                          text: 'Sing Out',
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
    });
  }
}
