//
//  UpdateSpotView.swift
//  WavySpots
//
//  Created by Alice Vedrenne on 18/05/2021.
//

import SwiftUI




//struct UpdateSpotView: View {
//    var spot: Spot
//
//    @State var showingAlert = false
//    var body: some View {
//        NavigationView {
//            var form = Form(titleForm: spot.Title, surfbreakForm: String, photoForm: <#T##String#>, levelForm: <#T##Int#>, addressForm: <#T##String#>)
//
//            Form {
//                Section {
//                    TextField("Destination", text: $titleForm)
//                    Picker(selection: $surfbreakForm,
//                           label: Text("Surf Break")) {
//                        ForEach(SurfBreak.allSurfBreaks, id: \.self) { surfBreak in
//                            Text(surfBreak).tag(surfBreak)
//                        }
//                    }
//                    TextField("Photo", text: $photoForm)
//
//                    Picker(selection: $levelForm,
//                           label: Text("Difficulty Level :")) {
//                        ForEach(Level.allLevels, id: \.self) { level in
//                            Text(String(level))
//                        }
//                    }
//                    TextField("Address", text: $addressForm)
//
//                }
////                Section {
////                    Button(action: {
////                        Api().addSpot(newSpot : Spot(ID: UUID().uuidString, Title: titleForm, Address: addressForm, Photo: photoForm, Level: levelForm, SurfBreak: [surfbreakForm]))
////                        print("Save Spot")
////                        showingAlert = true
////                    })
////                    {
////                        HStack {
////                            Text("Save Spot")
////                        }
////                        .alert(isPresented:$showingAlert) {
////                            Alert(
////                                title: Text("Your spot has been added")
////                                //primaryButton: .destructive(Text("Delete")) {
////                                //   print("Deleting...")
////                                //}
////                            )
////                        }
////                        //.frame(width: 50, height: 100, alignment: .center)
////                        .padding(10.0)
////                        .overlay(
////                            RoundedRectangle(cornerRadius: 10.0)
////                                .stroke(lineWidth: 2.0)
////                        )}
////                    //SaveSpot.center.x = self.view.center.x
////
////                }
////
////            }
////            .navigationBarTitle("Add Spot")
////
//
//        }
//}
//}
//
////struct UpdateSpotView_Previews: PreviewProvider {
////    static var previews: some View {
////        UpdateSpotView(spot: Spot())
////    }
////}
