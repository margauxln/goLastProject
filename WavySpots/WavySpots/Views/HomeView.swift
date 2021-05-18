import SwiftUI

struct HomeView: View {
    @State var isNavigationBarHidden: Bool = true

    @State var product: [Spot] = [Spot()]
    var body: some View {
        VStack(alignment: .leading) {
            HStack {
                CircleImage(photo:"logoSurf")
                VStack(alignment: .leading) {
                    Text("Wavy Spots")
                        .font(.title)
                        .fontWeight(.bold)
                        .foregroundColor(Color("Darkblue"))
                    Text("Best Places to surf ")
                        .foregroundColor(Color("BlueOcean"))
                        .font(Font.system(size:20, design: .default))
                }
            }
            .onAppear {
                Api().getSpots { spots in
                    product = spots
                    print(product)
                }
           }

            NavigationView {
               
                List {
                    ForEach(product, id: \.self) { Spot in
                        NavigationLink(
                            destination: DestinationPageView(spot:Spot)
                        ) {
                            BoxView(image: Spot.Photo, place: Spot.Address, description: Spot.SurfBreak[0])
                        }
                    }

                }
//                .onAppear {
//                    Api().getSpots { spots in
//                        product = spots
//                        print(product)
//                    }
//               }
                .navigationBarHidden(self.isNavigationBarHidden)
                            .onAppear {
                                self.isNavigationBarHidden = true
                            }
                
            }

        }
        .padding()
    }
}
//struct HomeView_Previews: PreviewProvider {
//    static var previews: some View {
//        HomeView()
//            .environment(\.sizeCategory, .extraSmall)
//            .previewDevice("iPhone 12 Pro Max")
//    }
//}
//

 
