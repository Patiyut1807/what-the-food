@file:Suppress("DEPRECATION")

package com.example.wtf_app

import android.app.Activity
import android.content.Intent
import android.net.Uri
import android.os.Bundle
import android.provider.MediaStore
import android.view.View
import androidx.activity.result.contract.ActivityResultContracts
import androidx.appcompat.app.AppCompatActivity
import com.android.volley.DefaultRetryPolicy
import com.android.volley.Response
import com.android.volley.toolbox.Volley
import com.example.wtf_app.databinding.ActivityMainBinding
import org.json.JSONObject
import java.io.IOException


const val REQUEST_CODE = 100


class MainActivity : AppCompatActivity() {
    private lateinit var viewBinding: ActivityMainBinding
    private var imageData: ByteArray? = null
    private val postURL: String = "http://154.215.14.243:8080/post-image"

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        viewBinding = ActivityMainBinding.inflate(layoutInflater)
        setContentView(viewBinding.root)

//        viewBinding.btnCamera.setOnClickListener { capturePhoto() }
        viewBinding.btnPick.setOnClickListener { openImageGallery() }
    }

//        private fun openCamera() {
//        val intent = Intent(
//            this,
//            CameraActivity::class.java
//        )
//        startActivity(intent)
//    }
    private fun capturePhoto() {

        val cameraIntent = Intent(MediaStore.ACTION_IMAGE_CAPTURE)
        ActivityResultContracts.RequestPermission()

    }

    private fun openImageGallery() {
        val intent = Intent(Intent.ACTION_PICK)
        intent.type = "image/*"
        startActivityForResult(intent, REQUEST_CODE)

//
    }
    private fun uploadImage() {
        imageData?: return
        val request = object : VolleyFileUploadRequest(
            Method.POST,
            postURL,
            Response.Listener {res->
                val jsonObject = JSONObject(String(res.data,charset("UTF-8")))
                stopProgressbar()
                viewBinding.textResult.text = "Result: ${jsonObject.getString("class") }"
                viewBinding.textProb.text = "Probability: %.2f".format(jsonObject.getDouble("probability"))
            },
            Response.ErrorListener {
                viewBinding.textResult.text = "Error"
            }
        ) {
            override fun getByteData(): MutableMap<String, FileDataPart> {
                var params = HashMap<String, FileDataPart>()
                params["image"] = FileDataPart("image", imageData!!, "jpg")
                return params
            }
        }
        request.setRetryPolicy(
            DefaultRetryPolicy(
                12000,
                DefaultRetryPolicy.DEFAULT_MAX_RETRIES,
                DefaultRetryPolicy.DEFAULT_BACKOFF_MULT
            )
        )
        Volley.newRequestQueue(this).add(request)
    }

    @Throws(IOException::class)
    private fun createImageData(uri: Uri) {
        val inputStream = contentResolver.openInputStream(uri)
        inputStream?.buffered()?.use {
            imageData = it.readBytes()
        }
    }

    override fun onActivityResult(requestCode: Int, resultCode: Int, data: Intent?) {
        if (resultCode == Activity.RESULT_OK && requestCode == REQUEST_CODE) {
            val uri = data?.data
            if (uri != null) {
                viewBinding.imgPreview.setImageURI(uri)
                startProgressbar()
                createImageData(uri)
                uploadImage()
            }
        }
        super.onActivityResult(requestCode, resultCode, data)
    }
    fun stopProgressbar(){
        viewBinding.progressBar.visibility = View.INVISIBLE
        viewBinding.textResult.visibility = View.VISIBLE
        viewBinding.textProb.visibility = View.VISIBLE
    }
    fun startProgressbar(){
        viewBinding.progressBar.visibility = View.VISIBLE
        viewBinding.textResult.visibility = View.INVISIBLE
        viewBinding.textProb.visibility = View.INVISIBLE
    }
    }




