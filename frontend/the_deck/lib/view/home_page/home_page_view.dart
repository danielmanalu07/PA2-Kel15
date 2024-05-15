import 'package:flutter/material.dart';
import 'package:the_deck/view/home_page/components/food_part.dart';
import 'package:the_deck/view/home_page/widgets/categories.dart';
import 'package:the_deck/view/home_page/widgets/discount_cart.dart';
import 'package:the_deck/view/home_page/widgets/popular.dart';
import 'package:the_deck/view/home_page/widgets/recommed.dart';
import 'package:the_deck/view/home_page/widgets/search_food.dart';
import 'package:the_deck/view/home_page/widgets/username_text.dart';
import 'components/size_config.dart';

class HomePageView extends StatefulWidget {
  const HomePageView({Key? key}) : super(key: key);

  @override
  _HomePageViewState createState() => _HomePageViewState();
}

class _HomePageViewState extends State<HomePageView> {
  @override
  Widget build(BuildContext context) {
    SizeConfig().init(context);
    return Scaffold(
      backgroundColor: Colors.white,
      body: SingleChildScrollView(
        child: Column(
          children: [
            UserNameText(),
            SearchFood(),
            DiscountCard(),
            FoodPart(partName: "Categories"),
            CategoriesFood(),
            FoodPart(partName: "Recommed"),
            RecommedFoods(),
            FoodPart(partName: "Popular"),
            PopularFoods(),
          ],
        ),
      ),
    );
  }
}
