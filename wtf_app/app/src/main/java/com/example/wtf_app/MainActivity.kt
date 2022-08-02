package com.example.wtf_app

import android.app.Activity
import android.content.Intent
import android.os.Bundle
import androidx.appcompat.app.AppCompatActivity
import com.example.wtf_app.databinding.ActivityMainBinding


const val REQUEST_CODE = 100

@Suppress("DEPRECATION")
class MainActivity : AppCompatActivity() {
    private lateinit var viewBinding: ActivityMainBinding

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        viewBinding = ActivityMainBinding.inflate(layoutInflater)
        setContentView(viewBinding.root)

        // Set up the listeners for take photo and video capture buttons
        viewBinding.btnCamera.setOnClickListener { openCamera() }
        viewBinding.btnPick.setOnClickListener { openImageGallery() }
    }

    private fun openCamera() {
        val intent = Intent(
            this,
            CameraActivity::class.java
        )
        startActivity(intent)
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
            viewBinding.imgPreview.setImageURI(data?.data) // handle chosen image
        }
    }

}