# Golang GCP Pubusb dengan Dead Letter Queue(DLQ)

> This edition of Google Cloud Compute (GCP)

Dalam Google Cloud Pub/Sub, DLQ (Dead Letter Queue) adalah mekanisme yang digunakan untuk menangani pesan yang gagal diproses oleh konsumen utama atau gagal diakui. DLQ memungkinkan pesan-pesan ini disimpan secara terpisah sehingga mereka dapat dianalisis atau diolah ulang nanti. Dalam kode Anda, DLQ digunakan untuk mengonsumsi pesan yang mungkin telah gagal dalam Subscriber utama dan diteruskan ke Subscriber DLQ (mocca-dlq-topic-sub) untuk pengolahan lebih lanjut atau pemantauan.

Berikut adalah beberapa poin terkait DLQ dan Pub/Sub dalam kode Anda:

Subscriber Utama (mocca-topic-sub):

Goroutine pertama (Worker Default) bertanggung jawab untuk mengkonsumsi pesan dari Subscriber utama (mocca-topic-sub).
Jika pesan yang diterima adalah "KIMINOTO" atau "EXIT", maka pesan tersebut ditandai sebagai pesan yang gagal diproses (msg.Nack()).
Pesan yang diterima yang bukan "KIMINOTO" atau "EXIT" akan diakui (msg.Ack()).
Subscriber DLQ (mocca-dlq-topic-sub):

Goroutine kedua (Worker DLQ) bertanggung jawab untuk mengkonsumsi pesan dari Subscriber DLQ (mocca-dlq-topic-sub).
Dalam kode yang Anda berikan, pesan yang diterima oleh Subscriber DLQ (mocca-dlq-topic-sub) diakui (msg.Ack()).
Ini berarti pesan yang mungkin telah gagal dalam Subscriber utama akan dipindahkan ke DLQ untuk pengolahan lebih lanjut atau pemantauan.
Pub/Sub Client (SetupClientGCPPubsub):

Fungsi SetupClientGCPPubsub digunakan untuk membuat klien Pub/Sub yang digunakan dalam kode Anda.
Klien ini digunakan untuk membuat dan mengelola Subscriber, serta mengonsumsi pesan dari topik Pub/Sub yang sesuai.
Dalam skenario penggunaan ini, DLQ berperan sebagai tempat penampungan untuk pesan-pesan yang mungkin perlu dianalisis atau diolah ulang jika mereka gagal dalam Subscriber utama. Penggunaan DLQ ini dapat membantu Anda memproses pesan yang gagal dengan lebih baik dan menghindari kehilangan data yang penting dalam sistem Anda.


Untuk Enable supaya DLQ Bisa berjalan dengan baik adalah sebagai berikut:
![plot](./assets/01.png)
Grant all Acccess Publisher role
![plot](./assets/02.png)
