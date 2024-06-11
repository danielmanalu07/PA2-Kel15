import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:the_deck/Controller/TableController.dart';
import 'package:the_deck/Core/app_colors.dart';
import 'package:the_deck/Core/font_size.dart';
import 'package:the_deck/Core/response_conf.dart';
import 'package:the_deck/Core/text_styles.dart';
import 'package:the_deck/Presentation/Table/RequestTableScreen.dart';

class Tablelistscreen extends StatefulWidget {
  const Tablelistscreen({Key? key}) : super(key: key);

  @override
  State<Tablelistscreen> createState() => _TablelistscreenState();
}

class _TablelistscreenState extends State<Tablelistscreen> {
  final TableController tableController = Get.put(TableController());

  @override
  void initState() {
    super.initState();
    tableController.fetchTables();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('All Tables'),
        actions: [
          IconButton(
            icon: Icon(Icons.request_page),
            onPressed: () => Get.to(
                () => RequestTableScreen()), // Navigate to RequestTableScreen
          ),
        ],
      ),
      body: Obx(() {
        if (tableController.isLoading.value) {
          return Center(child: CircularProgressIndicator());
        }
        return ListView.builder(
          itemCount: tableController.tables.length,
          itemBuilder: (context, index) {
            final table = tableController.tables[index];
            return Container(
              margin: EdgeInsets.all(8),
              padding: EdgeInsets.all(16),
              decoration: BoxDecoration(
                color: Pallete.neutral10,
                boxShadow: const [
                  BoxShadow(
                    color: Color(0x0A111111),
                    blurRadius: 24,
                    offset: Offset(0, 4),
                    spreadRadius: 0,
                  )
                ],
                borderRadius: BorderRadius.circular(16),
              ),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text(
                    'Table ${table.number}',
                    style: TextStyles.bodyMediumRegular.copyWith(
                        color: Pallete.neutral100,
                        fontSize: getFontSize(FontSizes.medium)),
                  ),
                  Text(
                    'Capacity: ${table.capacity}',
                    style: TextStyles.bodySmallRegular.copyWith(
                        color: Pallete.neutral70,
                        fontSize: getFontSize(FontSizes.small)),
                  ),
                  Text(
                    table.status == 1 ? 'Tidak Tersedia' : 'Tersedia',
                    style: TextStyles.bodySmallBold.copyWith(
                        color: table.status == 1 ? Colors.red : Colors.green,
                        fontSize: getFontSize(FontSizes.small)),
                  ),
                ],
              ),
            );
          },
        );
      }),
    );
  }
}
