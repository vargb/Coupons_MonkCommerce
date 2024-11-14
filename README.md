# Coupons_MonkCommerce

**Coupon Management**

    Implementations and Assumptions -
      1. Only a single coupon is applied at a time
      
      2. For each Coupon Name only a single type of coupon is attached, for example - consider coupon name "MONK" where only cart-wise coupon is attached and nothing else
      
      3. BxGy coupon when considered is only applied once when encountered. But for multiple occurences of this coupon have not implemented.
      
      4. Fixed prices throughout the Cart process.



    Code Shortcomings - 
      1. Have not handled duplicate coupon names
      2. Have not handled coupon/discount stacking i.e. handling multiple coupons at once, for example - Cart-wise coupons apply to entire cart, but certain products in the cart already have individual discounts or even other product or bxgy discounts. 
      3. Have not handled exceeding repetition limit or insufficient quantity
      4. Have not handled expired or inactive coupons.
      5. Consider both product-wise coupon and BxGy coupon is saving the same amount of monetary value to the customer but BxGy coupons gives free stuff that the customer might/might not need. So even in equal monetary savings and one coupon per cart rule, customer can be asked to choose between the types of coupons that they want to use.
      6. Would improve this by implementing auth-based user login and roles so everybody has different carts.
      7. Would improve BxGy coupon and product-wise coupon logic
      8. Would improve coupon precedence structure i.e. which coupon to be applied if multiple coupons are applicable

   Edge Cases or Potential Issues -
   
     1. In coupons the price set is static and the underlying product may increase/decrease in value which should also invariably increase/decrease the coupon discount percentage
     
     2. Real-time checking of inventory is not available and this may cause some issues in BxGy coupons.
     
     3. Variations in products might lead to some complications, for example consider product "A" as 2 variations small "a" and regular "A". Applying BxGy or product-wise coupons should be clearly assessed.
     
     4. Coupon stacking is also an issue, where some customers might exploit the system even if we restrict only one type of coupon per cart. So dialling this towards the customer or the client might be deliberated further. For example - consider BxGy coupon applied for products "A", "B" and "C". Customer buys all of em with applying cart-wise and product-wise coupons. So customer might get "D" product free and even pay low amount or nothing at all.
     
     5. Repetition Limit is set for BxGy but not for product-wise coupon.
     
     6. Coupons are to incentivize the customer(s) to spend more on the counter and also to get them back to the store for more in a later date. To set them accordingly is challenging as it paramount for the client not to incur too many losses as well as not to dissuade customers setting of coupons was not necessary. 
     

      
