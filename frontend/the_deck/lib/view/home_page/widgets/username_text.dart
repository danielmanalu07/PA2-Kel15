import 'package:flutter/material.dart';
import 'package:the_deck/view/home_page/components/colors.dart';
import 'package:the_deck/view/home_page/components/size_config.dart';

class UserNameText extends StatefulWidget {
  const UserNameText({Key? key}) : super(key: key);

  @override
  _UserNameTextState createState() => _UserNameTextState();
}

class _UserNameTextState extends State<UserNameText> {
  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: EdgeInsets.fromLTRB(
          SizeConfig.screenWidth!/14.17,                 /// 29.0
          SizeConfig.screenHeight!/23.55,                /// 29.0
          SizeConfig.screenWidth!/11.74,                 /// 35.0
          SizeConfig.screenHeight!/68.3                  /// 10.0
      ),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        crossAxisAlignment: CrossAxisAlignment.center,
        children: [
          Column(
            mainAxisAlignment: MainAxisAlignment.center,
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Padding(
                    padding: EdgeInsets.only(top: SizeConfig.screenHeight!/85.38 ,bottom: SizeConfig.screenHeight!/113.84),          /// 8.0 - 6.0
                    child: Text("Hi Kristina ", style: TextStyle(fontSize: SizeConfig.screenHeight!/29.7, fontWeight: FontWeight.bold, color: Colors.black87),),  /// 23
                  ),
                 
                ],
              ),
              Text("Are you hungry?", style: TextStyle(fontSize: SizeConfig.screenHeight!/40.18, color: Colors.black54),)   /// 17
            ],
          ),
          
        ],
      ),
    );
  }
}
