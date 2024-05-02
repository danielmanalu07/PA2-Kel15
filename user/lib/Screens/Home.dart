import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;

class Home extends StatelessWidget {
  const Home({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Home"),
      ),
      body: Center(
        child: ElevatedButton(
          onPressed: () {
            Navigator.push(
              context,
              MaterialPageRoute(
                builder: (context) => DataScreen(),
              ),
            );
          },
          child: Text("Get Data"),
        ),
      ),
    );
  }
}

class DataScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Data List"),
      ),
      body: FutureBuilder(
        future: fetchData(),
        builder: (context, snapshot) {
          if (snapshot.connectionState == ConnectionState.waiting) {
            return Center(child: CircularProgressIndicator());
          } else if (snapshot.hasError) {
            return Center(child: Text("Error: ${snapshot.error}"));
          } else {
            List<dynamic> data = snapshot.data as List<dynamic>;

            return ListView.builder(
              itemCount: data.length,
              itemBuilder: (context, index) {
                var item = data[index];
                return ListTile(
                  title: Text(item['name']),
                  subtitle: Text(item['description']),
                );
              },
            );
          }
        },
      ),
    );
  }
}

Future<List<dynamic>> fetchData() async {
  String url = "http://192.168.30.215:8002/category";

  final response = await http.get(Uri.parse(url));

  if (response.statusCode == 200) {
    final Map<String, dynamic> responseData = jsonDecode(response.body);
    List<dynamic> dataList = responseData['message'] ?? [];
    return dataList;
  } else {
    throw Exception('Failed to load data. Status code: ${response.statusCode}');
  }
}
