import 'package:flutter/material.dart';
import 'package:the_deck/view/home_page/components/colors.dart';
import 'package:the_deck/view/home_page/components/size_config.dart';

class ReviewStars extends StatefulWidget {
  const ReviewStars({Key? key}) : super(key: key);

  @override
  _ReviewStarsState createState() => _ReviewStarsState();
}

class _ReviewStarsState extends State<ReviewStars> {
  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: EdgeInsets.only(top: SizeConfig.screenHeight!/45.54),                /// 15.0
      
    );
  }
}
