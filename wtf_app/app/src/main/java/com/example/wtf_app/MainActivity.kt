@file:Suppress("DEPRECATION")

package com.example.wtf_app

import android.app.Activity
import android.app.ProgressDialog
import android.content.ActivityNotFoundException
import android.content.Intent
import android.graphics.Bitmap
import android.os.Bundle
import android.provider.MediaStore
import androidx.appcompat.app.AppCompatActivity
import com.android.volley.RequestQueue
import com.example.wtf_app.databinding.ActivityMainBinding


const val REQUEST_CODE = 100


class MainActivity : AppCompatActivity() {
    private lateinit var viewBinding: ActivityMainBinding


    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        viewBinding = ActivityMainBinding.inflate(layoutInflater)
        setContentView(viewBinding.root)

        viewBinding.btnCamera.setOnClickListener { capturePhoto() }
        viewBinding.btnPick.setOnClickListener { openImageGallery() }
    }

    //    private fun openCamera() {
//        val intent = Intent(
//            this,
//            CameraActivity::class.java
//        )
//        startActivity(intent)
//    }
    private fun capturePhoto() {

        val cameraIntent = Intent(MediaStore.ACTION_IMAGE_CAPTURE)
        try {
            startActivityForResult(cameraIntent, REQUEST_CODE)
        } catch (exception: ActivityNotFoundException) {
            //set this to request permistion
        }

    }

    private fun openImageGallery() {
        val intent = Intent(Intent.ACTION_PICK)
        intent.type = "image/*"
        startActivityForResult(intent, REQUEST_CODE)
    }



    @Deprecated("Deprecated in Java")
    override fun onActivityResult(requestCode: Int, resultCode: Int, data: Intent?) {
        super.onActivityResult(requestCode, resultCode, data)

        if (resultCode == Activity.RESULT_OK && requestCode == REQUEST_CODE) {
            viewBinding.imgPreview.setImageURI(data?.data)
        } else if (resultCode == Activity.RESULT_OK && requestCode == REQUEST_CODE && data != null) {
            viewBinding.imgPreview.setImageBitmap(data.extras?.get("data") as Bitmap)
        }

    }
}