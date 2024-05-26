import 'package:get/get.dart';
import 'package:the_deck/Controller/CustomerController.dart';
import 'package:the_deck/Core/app_colors.dart';
import 'package:the_deck/Core/assets_constantes.dart';
import 'package:the_deck/Core/response_conf.dart';
import 'package:the_deck/Presentation/Auth/screens/default_button.dart';
import 'package:the_deck/Presentation/Auth/screens/read_only.dart';
import 'package:the_deck/Presentation/Base/base.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:gap/gap.dart';

class PersonalDataView extends StatefulWidget {
  const PersonalDataView({Key? key}) : super(key: key);

  @override
  State<PersonalDataView> createState() => _PersonalDataViewState();
}

class _PersonalDataViewState extends State<PersonalDataView> {
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
            isBackup: false,
          ),
          body: Center(child: CircularProgressIndicator()),
        );
      }
      return Scaffold(
        appBar:
            buildAppBar(buildContext: context, screenTitle: "Personal Date"),
        body: Padding(
          padding: EdgeInsets.symmetric(horizontal: getWidth(24)),
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
                  ),
                ],
              ),
              const Gap(24),
              Column(
                children: [
                  ReadOnlyField(labelText: "Full Name", value: customer.name),
                  const Gap(12),
                  ReadOnlyField(
                      labelText: "Date of birth", value: customer.dateOfBirth),
                  const Gap(12),
                  ReadOnlyField(labelText: "Phone", value: customer.phone),
                  const Gap(12),
                  ReadOnlyField(labelText: "Email", value: customer.email),
                  const Gap(12),
                  ReadOnlyField(labelText: "Address", value: customer.address),
                  const Gap(12),
                  ReadOnlyField(labelText: "Gender", value: customer.gender),
                ],
              ),
              const Gap(36),
            ],
          ),
        ),
      );
    });
  }
}
