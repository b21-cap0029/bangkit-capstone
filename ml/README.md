# Dataset creation

We first fetched around 2700 entries at the beginning of May for tweets that contain certain keywords like depres, cemas, or gelisah. This is the __main dataset__ that will be annotated. After that, we fetched another 10000 random tweets that did not contain any of those keywords. These act as the “random noise” and are labeled as negative instances.

Three annotators annotated the main dataset using the guidelines mentioned in the reference paper. We then calculated the measure of agreement achieved between annotators using Krippendorf’s alpha, and we get the same result as the reference paper, which is around 0.6.

The main dataset was split into training and testing with 80%:20% ratio. This became __datd_train__ and __datd_test__. After that, some positive instances in the training dataset were combined with the random entries, creating a __random dataset__. This was meant to simulate random noise when grabbing and classifying tweets in production. This random dataset was also used for the final evaluation of the model.

# Model training and evaluation

Before training the model, we did some preprocessing on the dataset. This includes removing all links, emoticons, and strange symbols. 

After doing the preprocessing, we fed this into the ML model. The ML model uses an encoder to vectorize the sentence and a bidirectional LSTM layer for the training process. In the end, we managed to get a result that is comparable to the reference paper. We then save this model and pass this to the backend team.
