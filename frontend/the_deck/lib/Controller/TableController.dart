import 'package:get/get.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import 'package:the_deck/Models/Table.dart';

class TableController extends GetxController {
  var tables = <Table>[].obs;
  var isLoading = true.obs;

  @override
  void onInit() {
    fetchTables();
    super.onInit();
  }

  void fetchTables() async {
    try {
      isLoading(true);
      final response =
          await http.get(Uri.parse('http://192.168.66.215:8080/table'));
      if (response.statusCode == 200) {
        List<dynamic> data = jsonDecode(response.body)['message'];
        tables.value = data.map((json) => Table.fromJson(json)).toList();
      } else {
        Get.snackbar('Error', 'Failed to fetch tables');
      }
    } catch (e) {
      Get.snackbar('Error', 'Failed to fetch tables');
    } finally {
      isLoading(false);
    }
  }
}
