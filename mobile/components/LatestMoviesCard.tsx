import { StyleSheet, Text, View, Image, TouchableOpacity } from "react-native";
import React from "react";

const LatestMoviesCard = () => {
   return (
      <View className="mr-2 w-44 rounded-2xl flex items-center justify-center flex-col mx-1">
         <Image
            source={{
               uri: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR5sLgseJtBTleFdTJdNZ0crrvmIYNbHleJUTLBC1KkRQ&s",
            }}
            className="relative w-40 h-44  rounded-xl"
         />

         <Text className="text-brand absolute top-3 right-2 text-xs bg-brandLight rounded-full px-[6px] py-[1px]">
            Series
         </Text>

         <View>
            <Text className="text-light font-bold text-lg">
               Avatar: The Way of Water
            </Text>
            <Text className="text-mid font-bold text-md">3h 12m</Text>
         </View>
      </View>
   );
};

export default LatestMoviesCard;

const styles = StyleSheet.create({});