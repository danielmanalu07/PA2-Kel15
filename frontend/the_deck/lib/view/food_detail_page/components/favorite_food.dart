import 'package:flutter/material.dart';
import 'package:the_deck/view/home_page/components/size_config.dart';

class FavoriteFood extends StatefulWidget {
  const FavoriteFood({Key? key}) : super(key: key);

  @override
  _FavoriteFoodState createState() => _FavoriteFoodState();
}

class _FavoriteFoodState extends State<FavoriteFood> {
  @override
  Widget build(BuildContext context) {
    SizeConfig().init(context);
    return IconButton(
      onPressed: (){} ,
      icon: Icon(Icons.favorite),
      color: Colors.white,
      iconSize: SizeConfig.screenHeight!/22.77,
    );
  }
}
