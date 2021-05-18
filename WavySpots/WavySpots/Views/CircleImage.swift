//
//  surfImage.swift
//  WavySpots
//
//  Created by Fanny Armand on 03/05/2021.
//

import SwiftUI
import URLImage

struct CircleImage: View {
    var photo = ""
    var url = URL(string: "https://dl.airtable.com/ZuXJZ2NnTF40kCdBfTld_thomas-ashlock-64485-unsplash.jpg")!
    init(photo:String) {
        self.photo = photo
        if let url = URL(string: self.photo){
            self.url = url
        }
    }
    var body: some View {
        if photo == "logoSurf" {
            Image(photo)
                .resizable()
                .frame(width: 50, height: 50)
                .clipShape(Circle())
                .shadow(radius: 7)
        }
        else {
            URLImage(url: self.url,
                     content: { image in
                        image
                            .resizable()
                            .aspectRatio(contentMode: .fit)
                            .clipShape(Circle())
                     })
        }
            }
}

//struct CircleImage_Previews: PreviewProvider {
//    static var previews: some View {
//        CircleImage()
//    }
//}
